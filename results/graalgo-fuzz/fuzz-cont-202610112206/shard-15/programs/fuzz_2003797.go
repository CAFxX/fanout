package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int64 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint }
func genFn0(p0 string, p1 uint8, p2 int16) uint16 {
    v5 := []int{4294967295, 531}
    _ = len(v5)
    v6 := GenS1{F0: 104, F1: false, F2: 9223372036854775807}
    if (!(0 >= 0)) {
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 int64 = 42
            v9 := GenS1{F0: -1, F1: (false || ((p1 | 186) > 1)), F2: (v8 ^ (^(v8)))}
            _ = len(v5)
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    var v10 int8 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    _ = v6
    return 65535
}

func main() {
    v1 := []uint32{235, 2}
    var v2 uint32 = 79
    var v3 uint8 = 42
    var v4 uint8 = v3
    switch ("a" + "\n\t") {
    case "a":
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 bool = false
            var v13 int16 = -656
            _ = fi11
            _ = v12
            _ = v13
        }
    default:
        var v14 string = ("hello" + "a")
        _ = v14
    }
    v15 := GenS2{F0: -809, F1: -64}
    if true {
        var v16 int64 = -553
        _ = v16
    } else {
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int32 = 0
            _ = fi17
            _ = v18
        }
        v19 := GenS2{F0: -211, F1: 0}
        _ = v19
    }
    v4 = v3
    var v20 int64 = 951
    var v21 uint16 = 367
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 = v4
    switch int(v4) {
    case 32767:
        switch ((("\n\t" + "\n\t") + (("" + (("\n\t" + "hello") + (("hello" + "foo bar") + "a"))) + "a")) + "") {
        case "":
            var v22 float64 = (1e10 + 1.5)
            _ = v22
        case "hello":
            v1 = append(v1, (v2 >> ((993) & 63)))
        case "a":
            v23 := GenS3{F0: 255}
            _ = v23
        }
    case -9223372036854775808:
        v21 += 32767
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
