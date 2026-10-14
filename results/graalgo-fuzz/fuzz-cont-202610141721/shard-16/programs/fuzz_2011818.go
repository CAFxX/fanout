package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint; F2 uint8 }
func genFn0(p0 uint, p1 bool, p2 int16) int {
    fmt.Println(p2, p1)
    if p1 {
        var v5 int64 = -1
        var v6 int8 = int8(p0)
        v7 := GenS1{F0: 4294967295, F1: 611}
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 1, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2147483647
}

func main() {
    var v1 uint32 = 486
    v2 := []string{"a", (("foo bar" + ((("\n\t" + ("\n\t" + ("a" + "\n\t"))) + "\n\t") + "hello")) + (("hello" + "hello") + (("hello" + ("\n\t" + "foo bar")) + "hello")))}
    var v3 uint16 = 753
    v4 := make(map[int]string)
    if true {
        var v9 int64 = -1
        var v10 bool = (-2147483648 >= -1)
        if ((v10 || (1 == -1)) && v10) {
            if (1.5 > -2.25) {
                if ((v1 - v1) < 499) {
                    v11 := GenS1{F0: (v1 / ((v1) | 1)), F1: 42}
                    _ = v11
                }
                if true {
                    v10 = false
                }
            }
            fmt.Println(v10, v1, v9)
            for fi12 := 0; fi12 < 19; fi12++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v13 uint32 = (v1 >> ((2) & 63))
                _ = fi12
                _ = v13
            }
        }
        _ = v10
        _ = v9
    }
    v4[80] = ("" + "")
    v1 += v1
    var v14 uint32 = 2147483647
    switch ("hello" + "") {
    case "":
        var v15 float32 = 0.0
        _ = v15
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        fmt.Println(v3, fi16, v14)
        _ = fi16
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        if true {
            v18 := GenS1{F0: v1, F1: (v1 + v14)}
            var v19 int16 = 42
            v1 -= v14
            _ = v18
            _ = v19
        } else {
            var v20 int = 0
            v21 := v4[v20]
            var v22 string = "hello"
            _ = v20
            _ = v21
            _ = v22
        }
        v2 = append(v2, "a")
        _ = fi17
    }
    fmt.Println(v1)
    v14 -= (v1 / (((v1 % ((v14) | 1))) | 1))
    fmt.Println(v1, v1)
    if (false && (1 > 234)) {
        v23 := GenS1{F0: 0, F1: 219}
        switch "\n\t" {
        case "":
            fmt.Println(v14, v1)
        case "hello":
            _ = len(v4)
        }
        _ = v23
    }
    v1 = 2147483648
    for fi24 := 0; fi24 < 20; fi24++ {
        if false {
            _ = len(v4)
            fmt.Println(v14, v14)
        }
        v14 = v1
        _ = fi24
    }
    switch ("\n\t" + ("a" + "a")) {
    case "hello":
        v14 += v1
    case "foo bar":
        switch -292 {
        case 127:
            if true {
                for fi25 := 0; fi25 < 10; fi25++ {
                    var v26 int32 = 2147483647
                    v1 += 2147483648
                    _ = fi25
                    _ = v26
                }
                switch "hello" {
                case "hello":
                    v14 -= 65535
                }
                var v27 float32 = (0.1 * (1.5 + -2.25))
                _ = v27
            } else {
                v28 := GenS1{F0: (^(v14)), F1: 42}
                _ = v28
            }
        case 1:
            v29 := GenS2{F0: 127, F1: 42, F2: 128}
            _ = v29
        default:
            v3 += 1
        }
    case "":
        fmt.Println(v1)
    default:
        var v30 int32 = 42
        _ = v30
    }
    if len(v2) > 0 {
        v2[0] = ("hello" + "hello")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
