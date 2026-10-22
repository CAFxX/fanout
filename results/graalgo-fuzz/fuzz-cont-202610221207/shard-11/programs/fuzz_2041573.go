package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
func genFn0() string {
    if (2147483647 > -400) {
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 int8 = -83
            if (0 <= 42) {
                if true {
                    fmt.Println(v7, v7)
                    var v8 uint32 = 2
                    _ = v8
                }
            } else {
                var v9 uint8 = 1
                v10 := GenS1{F0: 32767}
                _ = v10
                _ = v9
            }
            _ = fi6
            _ = v7
        }
        var v11 string = "hello"
        v11 = (v11 + "")
        _ = v11
    }
    if true {
        var v12 uint = 24
        _ = v12
    } else {
        var v13 uint64 = 49
        var v14 string = "\n\t"
        v15 := []int{4294967295}
        _ = v13
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: -32768}
    _ = v16
    return "hello"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = -29
    var v3 uint = 835
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]bool)
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 bool = (0 >= 653)
        delete(v4, 505)
        _ = fi17
        _ = v18
    }
    v19 := []int{}
    v3 += v3
    v3 = 0
    for fi20 := 0; fi20 < 8; fi20++ {
        if true {
            for fi21 := 0; fi21 < 18; fi21++ {
                switch ("hello" + "foo bar") {
                case "":
                    _ = len(v1)
                case "hello":
                    var v22 int = v2
                    _ = v22
                case "\n\t":
                    v3 = (v3 * (v3 + v3))
                }
                v23 := GenS1{F0: 1}
                v3 -= (v3 | 435)
                _ = fi21
                _ = v23
            }
            var v24 string = "foo bar"
            fmt.Println(v2, v24, v2)
            _ = v24
        }
        _ = fi20
    }
    if false {
        v19 = append(v19, (v2 ^ v2))
        switch ("" + "hello") {
        case "hello":
            v3 += (v3 ^ (v3 * 834))
        default:
            var v25 bool = true
            _ = v25
        }
        if false {
            if (222 < 392) {
                _ = len(v5)
                v26 := GenS1{F0: -791}
                _ = v26
            }
            if false {
                var v27 int16 = 1000
                _ = v27
            } else {
                v28 := GenS2{F0: 4294967295}
                var v29 uint32 = 4294967295
                _ = v28
                _ = v29
            }
        } else {
            v3 += 65535
            if (2147483647 != -1) {
                for fi30 := 0; fi30 < 18; fi30++ {
                    var v31 float32 = ((0.0 / 1.0) + float32(0.1))
                    var v32 int64 = 1
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                v3 += v3
            } else {
                v3 = v3
            }
        }
    } else {
        _ = len(v19)
        fmt.Println(v2)
        fmt.Println(v3)
    }
    v3 = v3
    var v33 uint64 = 4294967295
    var v34 uint64 = (v33 | 2)
    switch "a" {
    case "hello":
        if (!(!true)) {
            fmt.Println(v33)
        }
    }
    for fi35 := 0; fi35 < 15; fi35++ {
        v2 = (v2 << (((v3 ^ v3)) & 63))
        for fi36 := 0; fi36 < 16; fi36++ {
            switch (fi35 / ((fi36) | 1)) {
            case 32767:
                var v37 string = ""
                _ = v37
            }
            var v38 float32 = (0.0 + 3.14159)
            if ((false && true) || false) {
                v33 = v34
            }
            _ = fi36
            _ = v38
        }
        _ = fi35
    }
    v39 := GenS1{F0: int16(v2)}
    fmt.Println(v34, v33, v33)
    v40 := GenS1{F0: 32767}
    var v41 int64 = 9223372036854775807
    var v42 uint64 = (v34 - 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
