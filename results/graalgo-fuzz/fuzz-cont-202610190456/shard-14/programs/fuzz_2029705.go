package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 int32, p1 uint32) uint32 {
    switch -2147483648 {
    case -1:
        v5 := GenS1{F0: 1}
        _ = v5
    case 4294967295:
        v6 := GenS2{F0: 42, F1: 32767, F2: 4294967295}
        _ = v6
    case 42:
        v7 := GenS2{F0: 255, F1: 0, F2: 2}
        _ = v7
    }
    switch ("hello" + ("\n\t" + "foo bar")) {
    case "\n\t":
        v8 := []int{}
        _ = v8
    case "foo bar":
        for fi9 := 0; fi9 < 5; fi9++ {
            p0 += -1
            _ = fi9
        }
    case "hello":
        if (true && true) {
            switch ("foo bar" + "\n\t") {
            case "hello":
                for fi10 := 0; fi10 < 6; fi10++ {
                    for fi11 := 0; fi11 < 4; fi11++ {
                        v12 := GenS3{F0: 42}
                        _ = fi11
                        _ = v12
                    }
                    var v13 float32 = 3.14159
                    var v14 int = int(p0)
                    _ = fi10
                    _ = v13
                    _ = v14
                }
            case "\n\t":
                switch ("a" + (("hello" + "hello") + "")) {
                case "":
                    if true {
                        v15 := []int{1, 9223372036854775807, -2147483648, 0}
                        _ = v15
                    }
                case "hello":
                    switch "\n\t" {
                    case "\n\t":
                        var v16 int = -2147483648
                        _ = v16
                    case "":
                        v17 := []int{4294967295, 32767, -884, -2147483648}
                        _ = v17
                    }
                case "a":
                    var v18 uint8 = 1
                    _ = v18
                }
            case "":
                var v19 int32 = p0
                _ = v19
            default:
                p0 = (^(p0))
            }
            p0 = p0
            v20 := []int{}
            _ = v20
        }
    default:
        p1 = (^(p1))
    }
    fmt.Println(p1, p0)
    p0 += p0
    v21 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = v21
    return (^(p1))
}

func main() {
    var v1 uint64 = 470
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    var v4 uint16 = 0
    v1 = v1
    v1 = (v1 / ((v1) | 1))
    v4 = v4
    var v22 int64 = 42
    v1 = v1
    var v23 int64 = v22
    switch 32768 {
    case 54:
        v4 = 759
    case -2147483648:
        var v24 int32 = -2147483648
        _ = v24
    }
    v4 += (^(v4))
    if (!true) {
        fmt.Println(v22, v4)
        fmt.Println(v22)
        if ((v22 % (((v23 * (v23 >> ((255) & 63)))) | 1)) > v23) {
            var v25 uint64 = (v1 << ((65535) & 63))
            if false {
                v26 := GenS3{F0: 1}
                _ = v26
            }
            _ = v25
        } else {
            fmt.Println(v23, v23, v23)
            v27 := GenS1{F0: 38}
            for fi28 := 0; fi28 < 19; fi28++ {
                for fi29 := 0; fi29 < 7; fi29++ {
                    v4 -= (^(v4))
                    var v30 uint = 255
                    _ = fi29
                    _ = v30
                }
                _ = fi28
            }
            _ = v27
        }
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        if true {
            if (true && true) {
                fmt.Println(v23, v23)
                switch "hello" {
                case "a":
                    fmt.Println(v23, v23)
                case "\n\t":
                    if ((((-(fi31)) != 9223372036854775807) && (v1 < (v1 >> ((18446744073709551615) & 63)))) || (!true)) {
                        delete(v2, "a")
                        v32 := GenS4{F0: 1, F1: v4, F2: -376}
                        _ = v32
                    } else {
                        var v33 float32 = (((1e10 + (0.0 - (0.1 / (-2.25 - 3.14159)))) / 3.14159) + (0.0 / 1e10))
                        _ = v33
                    }
                }
                if true {
                    v22 = 196
                    v22 -= 0
                } else {
                    if true {
                        v22 += (v23 - (v22 * v22))
                        var v34 string = "a"
                        v23 = 9223372036854775807
                        _ = v34
                    }
                    v35 := GenS4{F0: uint8(v1), F1: (v4 >> ((767) & 63)), F2: 2147483647}
                    v36 := GenS3{F0: -286}
                    _ = v35
                    _ = v36
                }
            }
            var v37 int8 = -1
            switch 32767 {
            case -259:
                v38 := GenS5{F0: (v22 | (v22 + 42)), F1: -2147483648}
                _ = v38
            case -2147483648:
                fmt.Println(v37, v22)
            default:
                fmt.Println(v23, v23, fi31)
            }
            _ = v37
        }
        _ = fi31
    }
    var v39 uint8 = 1
    var v40 int = 128
    delete(v2, "hello")
    for fi41 := 0; fi41 < 16; fi41++ {
        v42 := []int{v40, (fi41 | int(v39)), 954}
        v43 := GenS6{F0: 252}
        fmt.Println(fi41, v39)
        _ = fi41
        _ = v42
        _ = v43
    }
    var v44 uint = 18446744073709551615
    var v45 string = ""
    v46 := GenS3{F0: 220}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v44
    _ = v45
    _ = v46
}
