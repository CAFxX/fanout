package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int8; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint; F1 int }
func genFn0() uint64 {
    var v5 int8 = 0
    var v6 uint64 = 147
    _ = v5
    _ = v6
    return 809
}
func genFn1() uint {
    var v7 int16 = 42
    v8 := []int{32767, 4294967295}
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint32 = 253
    var v2 float64 = 1e10
    var v3 uint64 = 2
    var v4 int32 = 2147483647
    switch -260 {
    case 128:
        v9 := []int{32767, 32767, 127}
        _ = v9
    default:
        switch "foo bar" {
        case "hello":
            fmt.Println(v4, v4, v2)
        case "foo bar":
            fmt.Println(v1, v2)
        default:
            fmt.Println(v2, v1)
        }
    }
    if false {
        v1 += 65535
        var v10 uint32 = (^(v1))
        v11 := GenS1{F0: 2, F1: (("foo bar" + ("hello" + ("a" + (("" + "\n\t") + (("foo bar" + "\n\t") + ("foo bar" + (("" + ("foo bar" + ("hello" + (("" + "\n\t") + "hello")))) + "foo bar"))))))) + ("hello" + ""))}
        _ = v10
        _ = v11
    }
    var v12 int8 = 100
    for fi13 := 0; fi13 < 19; fi13++ {
        v1 -= 1
        switch 9223372036854775807 {
        case -179:
            v12 = v12
        case -775:
            var v14 int64 = 673
            _ = v14
        case 4294967295:
            v15 := []int{(fi13 ^ (fi13 ^ (fi13 | (fi13 * fi13)))), (fi13 - fi13)}
            _ = v15
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            if (-46 == (fi16 ^ (fi13 >> ((614) & 63)))) {
                fmt.Println(fi13, v2, v3)
            }
            v17 := GenS2{F0: (v12 ^ v12), F1: 146, F2: v1}
            _ = fi16
            _ = v17
        }
        _ = fi13
    }
    var v18 uint32 = 2147483648
    fmt.Println(v4, v3, v1)
    v12 -= v12
    var v19 uint64 = 0
    fmt.Println(v12)
    v20 := GenS3{F0: 447, F1: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
