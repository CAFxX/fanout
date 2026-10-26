package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int }
func genFn0(p0 int, p1 int32) uint8 {
    var v6 uint = 301
    v7 := GenS1{F0: 4294967295, F1: ""}
    if (true || true) {
        p1 += p1
    }
    var v8 int8 = 100
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 3.14159
    var v3 int8 = -47
    v4 := make(map[uint32]string)
    var v5 float32 = v2
    v4[42] = ("" + ("a" + "foo bar"))
    if true {
        switch "hello" {
        case "\n\t":
            for fi9 := 0; fi9 < 8; fi9++ {
                v3 = v3
                if false {
                    v10 := GenS2{F0: (fi9 % (((fi9 & fi9)) | 1))}
                    v11 := GenS1{F0: 4294967295, F1: "\n\t"}
                    _ = v10
                    _ = v11
                }
                _ = fi9
            }
        }
    }
    v12 := []int{-1, -9223372036854775808}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    if true {
        for fi13 := 0; fi13 < 19; fi13++ {
            v14 := GenS2{F0: (fi13 - (^(fi13)))}
            var v15 string = "foo bar"
            if (255 != 18446744073709551615) {
                var v16 uint64 = 650
                _ = v16
            } else {
                for fi17 := 0; fi17 < 4; fi17++ {
                    v2 += ((v5 - 1.5) * v2)
                    _ = fi17
                }
                v18 := GenS2{F0: (fi13 - fi13)}
                _ = v18
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        var v19 uint = 59
        _ = v19
    }
    if ((!(true && (!false))) || false) {
        v20 := GenS1{F0: 1, F1: ""}
        fmt.Println(v5, v2, v2)
        v1[42] = 244
        _ = v20
    } else {
        fmt.Println(v2, v2)
        v5 -= 3.14159
    }
    _ = len(v12)
    delete(v1, 2147483648)
    if false {
        var v21 string = "\n\t"
        v22 := GenS1{F0: 4294967295, F1: "hello"}
        _ = len(v4)
        _ = v21
        _ = v22
    } else {
        v5 -= -2.25
    }
    if (!false) {
        v2 = 0.1
        var v23 bool = ((!false) || (1 == 42))
        v2 = 3.14159
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
