package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int16, p1 int64, p2 uint) int64 {
    fmt.Println(p2, p0, p2)
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return (p1 + p1)
}
func genFn1(p0 int64, p1 int) int16 {
    var v4 bool = true
    var v5 uint64 = 687
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 float32 = ((-2.25 + 1e10) * 0.1)
    var v2 string = (("foo bar" + "") + ("\n\t" + ("hello" + "")))
    var v3 string = "a"
    if ((v3 + v2) == ((v2 + v2) + "a")) {
        v3 = v3
        var v6 uint16 = 0
        switch v2 {
        case "":
            switch v3 {
            case "hello":
                var v7 uint16 = (^(v6))
                _ = v7
            default:
                if false {
                    v8 := []int{74, -1}
                    _ = v8
                } else {
                    var v9 uint16 = (^(v6))
                    var v10 uint = 2
                    fmt.Println(v3, v6)
                    _ = v10
                    _ = v9
                }
            }
        case "hello":
            v1 = float32(0.1)
        }
        _ = v6
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS1{F0: v2, F1: ((v2 + v3) + v3)}
        var v13 int64 = 1
        var v14 uint = 66
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    if (210 != 42) {
        var v15 float32 = v1
        _ = v15
    }
    if true {
        fmt.Println(v2, v1)
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        if (true && (135 < 2)) {
            if (v1 < v1) {
                v17 := GenS2{F0: (-32768 > genFn1(1, (fi16 + 0))), F1: 255}
                _ = v17
            }
        } else {
            fmt.Println(fi16, v1, v2)
            var v18 int64 = 9223372036854775807
            _ = v18
        }
        v19 := []int{(fi16 - (fi16 | (fi16 << ((617) & 63)))), fi16, fi16}
        _ = fi16
        _ = v19
    }
    fmt.Println(v3)
    var v20 uint64 = 1
    if true {
        var v21 uint8 = 0
        v3 = v3
        _ = v21
    }
    var v22 int16 = genFn1(42, -875)
    var v23 uint64 = v20
    if true {
        var v24 string = ""
        v25 := GenS3{F0: 2147483648}
        _ = v24
        _ = v25
    }
    fmt.Println(v1, v23)
    v2 = v3
    v26 := []int{890, 9223372036854775807, -167}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v26
    _ = v3
}
