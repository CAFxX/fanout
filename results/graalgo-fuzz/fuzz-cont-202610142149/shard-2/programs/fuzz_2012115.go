package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 bool }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint }
func genFn0() uint8 {
    v5 := GenS1{F0: ("\n\t" + (("a" + "hello") + (("foo bar" + "") + "foo bar"))), F1: 1, F2: (-2147483648 >= 0)}
    var v6 string = ""
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int8) int64 {
    fmt.Println(p0)
    var v7 string = (("" + "a") + ((((("\n\t" + "\n\t") + "\n\t") + ("a" + "")) + "a") + ""))
    _ = p0
    _ = v7
    return -526
}

func main() {
    v1 := make(map[string]uint32)
    var v2 string = ""
    v3 := make(map[int]string)
    v4 := []uint32{782, 0, 0, 684}
    v8 := GenS1{F0: ((("" + (v2 + v2)) + "\n\t") + ((v2 + v2) + "hello")), F1: genFn1(-1), F2: (669 != genFn1(100))}
    switch (v2 + ("" + v2)) {
    case "":
        fmt.Println(v2)
    default:
        var v9 int = 127
        _ = v9
    }
    v2 = ""
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS2{F0: 1, F1: 4294967295}
        if true {
            if (1 <= 0) {
                v2 = ((v2 + v2) + "a")
                fmt.Println(fi10, v2)
            } else {
                if len(v4) > 0 {
                    v4[0] = 4294967295
                }
            }
            fmt.Println(fi10, v2)
        }
        v12 := GenS3{F0: 424}
        _ = fi10
        _ = v11
        _ = v12
    }
    v2 = v2
    _ = len(v3)
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 1; fi14++ {
            v2 = v2
            v2 = "hello"
            if true {
                var v15 uint64 = 861
                if (-199 >= -1) {
                    v15 -= (v15 + v15)
                }
                v15 -= v15
                _ = v15
            }
            _ = fi14
        }
        _ = fi13
    }
    _ = len(v3)
    if (32768 != 2) {
        if false {
            var v16 int8 = -80
            _ = v16
        } else {
            fmt.Println(v2)
            switch -900 {
            case 2147483647:
                fmt.Println(v2)
            case 101:
                var v17 int = -9223372036854775808
                _ = v17
            }
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v18 float64 = -2.25
    v2 = "\n\t"
    var v19 int64 = 918
    if true {
        var v20 int8 = 1
        _ = v20
    } else {
        v18 -= 0.0
    }
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
