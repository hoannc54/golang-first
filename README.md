# golang-first
Project đầu tiên về golang

# Golang

## Tài liệu
* Document: https://golang.org/doc/
* Các Packages : https://golang.org/pkg/
* Github: 
    - https://github.com/golang/go/wiki
    - 
* https://www.tutorialspoint.com/go/index.htm

### Book
* [An Introduction to Programming in Go](http://www.golang-book.com/books/intro)

## IDE: 
* Sublime Text: Package https://github.com/DisposaBoy/GoSublime

## Go tốt nhất cho việc gì?

Phân phối các network service (dịch vụ mạng). Các chương trình ứng dụng mạng (network application) sống hay chết là dựa vào concurrency và các tính năng native concurrency của Go, các goroutines và các channel, rất phù hợp cho các tác vụ đó. Do đó, có nhiều dự án Go dành cho mạng, các chức năng distributed (phân phối) và dịch vụ đám mây: API, web server, minimal frameworks cho các web application và các loại tương tự.

Sự phát triển của cloud-native. Các tính năng concurrency và network của Go và tính linh hoạt cao của nó làm cho nó phù hợp với việc xây dựng các ứng dụng cloud-native. Trên thực tế, Go đã được sử dụng để xây dựng một trong những nền tảng phát triển ứng dụng dựa trên cloud-native, ứng dụng hệ thống containerization Docker.

Thay thế cho cơ sở hạ tầng hiện có. Phần lớn các phần mềm của chúng tôi phụ thuộc vào cơ sở hạ tầng Internet đã lạc hậu. Việc viết lại những thứ như vậy bằng Go mang lại nhiều lợi ích, như giữ an toàn bộ nhớ tốt hơn, triển khai trên nhiều nền tảng dễ dàng hơn và một code base “sạch” để hỗ trợ bảo trì trong tương lai. Một server SSH mới được gọi là Teleport và một phiên bản mới của Network Time Protocol được viết bằng Go, được cung cấp như phương pháp thay thế cho các đối tác thông thường của họ.

## Command 

```
Go is a tool for managing Go source code.

Cách dùng:

    go command [arguments]

The commands are:

    build       compile packages and dependencies
    clean       remove object files and cached files
    doc         show documentation for package or symbol
    env         print Go environment information
    bug         start a bug report
    fix         update packages to use new APIs
    fmt         gofmt (reformat) package sources
    generate    generate Go files by processing source
    get         download and install packages and dependencies
    install     compile and install packages and dependencies
    list        list packages
    run         compile and run Go program
    test        test packages
    tool        run specified go tool
    version     print Go version
    vet         report likely mistakes in packages

Use "go help [command]" for more information about a command.

Additional help topics:

    c           calling between Go and C
    buildmode   build modes
    cache       build and test caching
    filetype    file types
    gopath      GOPATH environment variable
    environment environment variables
    importpath  import path syntax
    packages    package lists
    testflag    testing flags
    testfunc    testing functions

Use "go help [topic]" for more information about that topic.
```

## Go tour

Tour of go online : https://tour.golang.org/welcome/1
 
Go tour from Git: https://github.com/golang/tour

### Basic

#### Khai báo biến 
#### Gọi hàm 
#### Packages

#### for, if, else, switch, defer

**for**

**if**

**if with a short statement**

Viết thêm câu lệnh vào trước điều kiện
    
```
if v := math.Pow(x, n); v < lim {
        return v
    }
```
  

**if else**
* struct, slices, map

**defer**
Giữ hàm cho đến khi các hàm khác thực thi xong

**Pointer**

**Stacking defers**

Nếu có nhiều hơn 2 defer thì defer tuân theo luật stack để chạy 

### Method, interface

### Concurrency
