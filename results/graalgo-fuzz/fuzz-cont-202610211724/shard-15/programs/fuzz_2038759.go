package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int16) int64 {
    p0 = (p0 * p0)
    var v4 uint16 = 42
    var v5 string = ("a" + ("hello" + (("" + "") + (("foo bar" + "hello") + ("hello" + "foo bar")))))
    if true {
        switch "\n\t" {
        case "hello":
            var v6 int64 = -229
            _ = v6
        case "foo bar":
            v5 = "foo bar"
        default:
            var v7 int32 = 42
            _ = v7
        }
        fmt.Println(v5, v4, p0)
    } else {
        v4 += v4
        v8 := []int{-9223372036854775808, 4294967295, 682, 0}
        v9 := GenS1{F0: 629}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v4
    _ = v5
    return -9223372036854775808
}
func genFn1(p0 uint32, p1 uint16, p2 string) int8 {
    p1 -= p1
    switch "hello" {
    case "\n\t":
        if false {
            var v10 int = 32768
            var v11 int32 = 0
            _ = v10
            _ = v11
        } else {
            v12 := GenS1{F0: (^(p1))}
            switch 4294967295 {
            case 42:
                var v13 int64 = -9223372036854775808
                _ = v13
            case -9223372036854775808:
                p2 = (((("\n\t" + "foo bar") + (p2 + "foo bar")) + ("hello" + "")) + "\n\t")
            default:
                if ((p0 * (p0 & 65535)) <= 0) {
                    v14 := GenS1{F0: p1}
                    _ = v14
                } else {
                    p0 = (p0 << ((1) & 63))
                    v15 := GenS1{F0: 955}
                    _ = v15
                }
            }
            var v16 int32 = 2147483647
            _ = v12
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        if true {
            p1 -= p1
        }
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    var v1 int32 = 1
    var v2 float32 = (1e10 + 0.1)
    var v3 int16 = 0
    var v18 int8 = genFn1(42, 32767, "foo bar")
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 uint16 = 1
        _ = fi19
        _ = v20
    }
    v3 -= v3
    v21 := []int{int(v18), -807, 347}
    if false {
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
        if len(v21) > 0 {
            v21[0] = 906
        }
        var v22 uint16 = 42
        _ = v22
    }
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    v18 -= 127
    if false {
        v3 = (v3 & v3)
        v2 = -2.25
    }
    var v23 float32 = ((1.5 / v2) - float32(1e10))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v23
    _ = v3
}
