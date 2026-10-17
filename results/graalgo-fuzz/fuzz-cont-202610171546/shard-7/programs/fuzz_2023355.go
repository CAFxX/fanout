package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 uint32 }
func genFn0(p0 int32, p1 uint8) int8 {
    switch (("a" + "a") + "hello") {
    case "":
        fmt.Println(p1, p0)
    default:
        switch 2147483647 {
        case 118:
            var v5 bool = false
            _ = v5
        case -2147483648:
            v6 := []int{95, -778}
            _ = v6
        case 127:
            var v7 bool = true
            _ = v7
        }
    }
    switch 9223372036854775807 {
    case 128:
        p0 -= 717
    }
    _ = p0
    _ = p1
    return -69
}
func genFn1() int {
    if false {
        var v8 int64 = 0
        v8 = 0
        _ = v8
    }
    if ((("hello" + "a") + "a") < ("hello" + "a")) {
        var v9 int32 = -431
        v9 = 309
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := GenS1{F0: 1}
            v13 := []int{9223372036854775807, 42, 32767}
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    var v14 bool = true
    _ = v14
    return -2147483648
}

func main() {
    v1 := []string{("a" + ""), "", ("hello" + ""), ("a" + "")}
    var v2 uint16 = 772
    var v3 float32 = (0.1 * float32(float64(3.14159)))
    var v4 uint32 = 2147483647
    var v15 bool = (false && (!(true || false)))
    var v16 int64 = -1
    fmt.Println(v4, v15)
    if v15 {
        v17 := GenS1{F0: 947}
        _ = len(v1)
        _ = v17
    } else {
        v3 += (((v3 + v3) / (((v3 - (v3 - v3)) + (v3 - 0.0)) / (-2.25 - float32((1.5 / float64(v3)))))) - (v3 + v3))
        fmt.Println(v15)
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        v15 = (255 < v2)
        _ = fi18
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        switch (fi19 - (fi19 % (((fi19 + fi19)) | 1))) {
        case -2147483648:
            v2 += (v2 / ((v2) | 1))
        case -9223372036854775808:
            v3 = (0.0 * ((float32((((((0.0 * 0.1) - 1.5) / -2.25) + 1.5) + (1.5 / 1.5))) / ((((v3 / 1e10) - ((((float32(0.1) - (v3 + v3)) / -2.25) / 3.14159) / v3)) / ((v3 + v3) - v3)) - (3.14159 - (v3 - v3)))) * (1e10 + 1.5)))
        default:
            _ = len(v1)
        }
        if true {
            v2 = v2
            fmt.Println(v4, v4, v3)
            var v20 int = fi19
            _ = v20
        } else {
            if ((v4 | v4) == (v4 % (((v4 + 2147483647)) | 1))) {
                var v21 float64 = 1e10
                fmt.Println(v21)
                switch 528 {
                case 32768:
                    switch (-(fi19)) {
                    case -9223372036854775808:
                        var v22 float32 = (float32(((v21 - float64((v3 * v3))) * float64(v3))) - v3)
                        _ = v22
                    }
                }
                _ = v21
            } else {
                var v23 bool = (-91 < 0)
                var v24 int = -258
                _ = v23
                _ = v24
            }
        }
        var v25 float64 = (float64(((v3 / v3) / 0.1)) - 1.5)
        _ = fi19
        _ = v25
    }
    if (-1 != 487) {
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        v1 = append(v1, (("foo bar" + "hello") + "a"))
        _ = len(v1)
    } else {
        for fi26 := 0; fi26 < 15; fi26++ {
            var v27 int = -9223372036854775808
            v28 := GenS2{F0: 18446744073709551615}
            _ = fi26
            _ = v27
            _ = v28
        }
        switch 449 {
        case 32767:
            switch 128 {
            case 2147483647:
                var v29 float32 = -2.25
                _ = v29
            default:
                switch ("\n\t" + "foo bar") {
                case "\n\t":
                    v4 += (v4 >> ((2) & 63))
                case "a":
                    var v30 int16 = 42
                    _ = v30
                case "":
                    v31 := GenS3{F0: 91, F1: v4}
                    _ = v31
                }
            }
        case -892:
            if len(v1) > 0 {
                v1[0] = ""
            }
        }
        switch "\n\t" {
        case "\n\t":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            v4 = 2
        }
    }
    if (((1.5 - (1.5 / float64((v3 + v3)))) + (((1e10 + 1e10) + 1.5) * (1e10 - 1e10))) >= (0.1 * (float64(1e10) * (-2.25 / (1.5 * -2.25))))) {
        fmt.Println(v16, v16)
    } else {
        var v32 int64 = v16
        _ = v32
    }
    v15 = v15
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
