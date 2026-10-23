package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint; F2 uint }
type GenS3 struct { F0 int16; F1 uint16; F2 int }
func genFn0(p0 int32) uint32 {
    fmt.Println(p0)
    p0 = p0
    if (32767 > 32768) {
        var v4 string = ("\n\t" + "foo bar")
        p0 = p0
        p0 = 913
        _ = v4
    } else {
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 uint16 = 1
            v7 := []int{(fi5 + fi5), 872, (fi5 >> ((1) & 63))}
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := []int{-877}
        _ = v8
    }
    v9 := GenS1{F0: 42}
    _ = p0
    _ = v9
    return 384
}
func genFn1(p0 int32, p1 int8, p2 string) uint64 {
    if false {
        v10 := GenS1{F0: 42}
        fmt.Println(p1, p0)
        _ = v10
    }
    switch -9223372036854775808 {
    case 128:
        v11 := []int{420}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 692
}

func main() {
    var v1 int8 = 1
    var v2 int16 = 1
    var v3 int8 = v1
    var v12 uint16 = 0
    var v13 uint16 = (v12 % ((429) | 1))
    var v14 bool = (!false)
    v15 := GenS1{F0: v12}
    if (!v14) {
        var v16 int = 0
        _ = v16
    }
    var v17 float64 = -2.25
    fmt.Println(v1, v12, v17)
    fmt.Println(v13)
    v18 := GenS2{F0: v2, F1: 4294967295, F2: 237}
    if v14 {
        for fi19 := 0; fi19 < 3; fi19++ {
            v20 := GenS3{F0: v2, F1: 255, F2: fi19}
            fmt.Println(v12, fi19, v2)
            _ = fi19
            _ = v20
        }
    } else {
        v21 := []int{}
        v22 := GenS1{F0: (v13 / (((v12 * v13)) | 1))}
        _ = v21
        _ = v22
    }
    if (v14 || (((v2 - 0) > -78) && false)) {
        v13 = 32767
        fmt.Println(v12, v3, v17)
        v23 := []int{32768}
        _ = v23
    }
    switch "a" {
    case "a":
        switch ("" + "") {
        case "hello":
            v24 := []int{-618, 9223372036854775807, 9223372036854775807}
            _ = v24
        case "\n\t":
            v17 = (-2.25 + 3.14159)
        default:
            var v25 uint64 = 480
            _ = v25
        }
    default:
        var v26 float64 = (1.5 + ((-2.25 / float64(3.14159)) - ((v17 / ((v17 * v17) * 1e10)) / v17)))
        _ = v26
    }
    switch -250 {
    case -116:
        if (3.14159 == (1e10 * 0.0)) {
            v27 := []int{566, 32767, 1, -498}
            v28 := GenS1{F0: v12}
            _ = v27
            _ = v28
        }
    case 0:
        v29 := []int{-507}
        _ = v29
    case -9223372036854775808:
        v1 -= v3
    }
    if ((((!true) && (((!(1 >= 42)) && v14) && (-2147483648 <= -209))) && (!((v1 / ((v1) | 1)) == 1))) && (!(true || v14))) {
        v30 := GenS3{F0: -32768, F1: 32767, F2: -9223372036854775808}
        var v31 uint32 = 894
        fmt.Println(v13, v12, v2)
        _ = v30
        _ = v31
    } else {
        if v14 {
            var v32 float64 = float64((1.5 - -2.25))
            _ = v32
        }
        fmt.Println(v2, v17)
    }
    var v33 uint = 797
    v34 := GenS3{F0: (v2 * v2), F1: (v13 + v12), F2: int(v33)}
    if false {
        if v14 {
            v33 = (^(v33))
        }
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
}
