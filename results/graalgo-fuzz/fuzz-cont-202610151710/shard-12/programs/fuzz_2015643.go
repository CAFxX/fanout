package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() uint16 {
    var v4 int64 = 9223372036854775807
    var v5 float32 = 1.5
    var v6 uint = 2
    var v7 uint8 = 255
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 912
}
func genFn1(p0 uint8) int16 {
    var v8 float32 = float32(((3.14159 * 1.5) + 3.14159))
    if (255 != 4294967295) {
        var v9 uint32 = 420
        _ = v9
    }
    fmt.Println(p0, p0)
    if true {
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 int8 = -114
            var v12 uint = 833
            switch ("hello" + "hello") {
            case "\n\t":
                var v13 int16 = 1000
                _ = v13
            case "foo bar":
                for fi14 := 0; fi14 < 14; fi14++ {
                    var v15 string = "a"
                    v16 := []int{(fi14 - (fi14 ^ -377)), (fi14 & fi10), 0}
                    v8 -= v8
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case "":
                v17 := GenS1{F0: (v12 + (v12 + (v12 - v12)))}
                _ = v17
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    } else {
        v8 = 3.14159
        var v18 uint = 18446744073709551615
        p0 = 127
        _ = v18
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        fmt.Println(v8)
        _ = fi19
    }
    _ = p0
    _ = v8
    return 42
}

func main() {
    v1 := []int{}
    var v2 uint16 = 418
    v3 := make(map[int]int)
    var v20 bool = true
    switch "a" {
    case "a":
        if (v20 || (!(false || v20))) {
            delete(v3, 127)
            if false {
                v2 = (v2 << ((1) & 63))
            }
        } else {
            delete(v3, 980)
            if len(v1) > 0 {
                v1[0] = 1
            }
        }
    case "foo bar":
        var v21 float64 = 0.0
        _ = v21
    case "":
        if (625 >= 0) {
            for fi22 := 0; fi22 < 13; fi22++ {
                v20 = v20
                _ = fi22
            }
            var v23 string = "a"
            v23 = v23
            _ = v23
        } else {
            for fi24 := 0; fi24 < 8; fi24++ {
                switch "foo bar" {
                case "\n\t":
                    delete(v3, (fi24 % ((fi24) | 1)))
                default:
                    v20 = (-681 == 0)
                }
                if len(v1) > 0 {
                    v1[0] = fi24
                }
                _ = fi24
            }
            var v25 uint32 = 42
            _ = v25
        }
    default:
        if (1e10 != 3.14159) {
            fmt.Println(v20, v20)
            var v26 uint8 = 46
            v2 += (v2 / ((v2) | 1))
            _ = v26
        } else {
            var v27 int64 = -1
            _ = v27
        }
    }
    switch (("" + "") + ("foo bar" + ("" + ((("hello" + "") + "hello") + ("hello" + ""))))) {
    case "\n\t":
        v20 = false
    case "a":
        switch 42 {
        case 2147483647:
            fmt.Println(v2)
        }
    case "":
        var v28 string = "a"
        _ = v28
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        fmt.Println(fi29)
        var v30 float32 = 3.14159
        v30 -= float32(3.14159)
        _ = fi29
        _ = v30
    }
    fmt.Println(v2, v2)
    fmt.Println(v20, v20)
    v20 = v20
    v31 := v3[42]
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    var v32 uint16 = (v2 << ((810) & 63))
    v32 = (^(v32))
    var v33 uint = 444
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
