package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]string)
    var v4 string = "foo bar"
    v1[391] = 32768
    var v5 string = v4
    v4 = "\n\t"
    v6 := GenS1{F0: 2}
    var v7 uint64 = 1
    var v8 int16 = -718
    switch -23 {
    case -1:
        for fi9 := 0; fi9 < 20; fi9++ {
            fmt.Println(fi9, fi9)
            var v10 uint = 65535
            _ = fi9
            _ = v10
        }
    case 790:
        v7 += 0
    case 434:
        v11 := GenS2{F0: 466}
        _ = v11
    }
    _ = len(v1)
    if (694 <= 1) {
        var v12 uint8 = 42
        _ = v12
    } else {
        _ = len(v3)
        _ = len(v3)
        fmt.Println(v4, v2)
    }
    if false {
        v4 = (("hello" + v4) + v5)
        delete(v1, 640)
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        var v13 uint = 1
        _ = v13
    case 2147483647:
        if false {
            switch 4294967295 {
            case -425:
                fmt.Println(v5)
            case -1:
                v14 := GenS1{F0: 676}
                _ = v14
            }
            v5 = "a"
        } else {
            switch "hello" {
            case "foo bar":
                v15 := GenS3{F0: 357}
                _ = v15
            case "hello":
                fmt.Println(v2, v5, v2)
            default:
                var v16 int32 = -2147483648
                _ = v16
            }
            v5 = (v5 + v4)
            v5 = (v5 + v5)
        }
    }
    switch -1 {
    case 0:
        var v17 uint8 = 39
        _ = v17
    case -2147483648:
        var v18 float32 = (1e10 / 1.5)
        _ = v18
    case 9223372036854775807:
        v19 := GenS4{F0: 398}
        _ = v19
    }
    switch -15 {
    case 42:
        v5 = "\n\t"
    case -1:
        v20 := GenS1{F0: 267}
        _ = v20
    case -137:
        v5 = ((v4 + ("foo bar" + "\n\t")) + "\n\t")
    default:
        switch 1 {
        case -796:
            v21 := GenS1{F0: 2}
            _ = v21
        }
    }
    var v22 int64 = (v2 * 9223372036854775807)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
