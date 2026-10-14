package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint16 }
func genFn0() string {
    var v5 uint64 = 2
    v6 := GenS1{F0: 65535}
    var v7 uint64 = (^(v5))
    switch 0 {
    case 0:
        if true {
            if false {
                var v8 int8 = 42
                _ = v8
            }
        }
    case -2147483648:
        v9 := GenS2{F0: 255, F1: v7}
        _ = v9
    }
    _ = v5
    _ = v6
    _ = v7
    return (("hello" + ((("" + "foo bar") + "foo bar") + "hello")) + "hello")
}

func main() {
    var v1 float64 = (3.14159 / ((((0.0 / (1.5 - float64(3.14159))) + float64(3.14159)) + float64(0.0)) / 1.0))
    var v2 uint8 = 255
    var v3 int8 = 42
    var v4 int32 = -1
    var v10 int16 = -880
    v11 := GenS3{F0: 191, F1: 255}
    v10 -= (v10 >> ((18446744073709551615) & 63))
    var v12 int = -9223372036854775808
    if true {
        var v13 uint64 = 0
        _ = v13
    }
    var v14 float64 = 3.14159
    for fi15 := 0; fi15 < 5; fi15++ {
        fmt.Println(v3, v12)
        _ = fi15
    }
    var v16 string = "hello"
    if ((v10 * (v10 | v10)) < v10) {
        if false {
            if false {
                var v17 bool = true
                var v18 uint64 = 55
                _ = v17
                _ = v18
            }
        }
        v19 := GenS3{F0: 42, F1: 1}
        _ = v19
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        for fi21 := 0; fi21 < 10; fi21++ {
            if ((4294967295 >= 42) || true) {
                v10 += (v10 & v10)
                var v22 string = (v16 + ("\n\t" + v16))
                v2 = (v2 % ((v2) | 1))
                _ = v22
            } else {
                if false {
                    var v23 float64 = -2.25
                    var v24 uint16 = 255
                    var v25 uint16 = (v24 - v24)
                    _ = v23
                    _ = v24
                    _ = v25
                } else {
                    fmt.Println(v14, fi21, v12)
                }
                var v26 int = (fi21 << ((65535) & 63))
                v27 := GenS1{F0: 0}
                _ = v26
                _ = v27
            }
            fmt.Println(v2)
            var v28 bool = false
            _ = fi21
            _ = v28
        }
        v1 = v14
        fmt.Println(v14, fi20, v10)
        _ = fi20
    }
    switch v16 {
    case "hello":
        v4 -= (v4 & v4)
    }
    fmt.Println(v14)
    switch (v12 / (((^(v12))) | 1)) {
    case -507:
        fmt.Println(v3, v4)
    case 9223372036854775807:
        var v29 float32 = (0.0 + 3.14159)
        _ = v29
    default:
        fmt.Println(v10, v10)
    }
    if true {
        if (((v2 << ((255) & 63)) == v2) && false) {
            var v30 int16 = v10
            for fi31 := 0; fi31 < 2; fi31++ {
                var v32 float64 = (v1 / (v1 + 1.5))
                var v33 uint = 255
                _ = fi31
                _ = v32
                _ = v33
            }
            _ = v30
        } else {
            v4 = 1
            for fi34 := 0; fi34 < 18; fi34++ {
                var v35 float32 = 1e10
                v14 = 1.5
                _ = fi34
                _ = v35
            }
            v36 := []int{v12}
            _ = v36
        }
        if ((!true) || false) {
            var v37 uint32 = 995
            v14 += v14
            if true {
                var v38 string = ("foo bar" + "")
                v1 -= float64(float32(v1))
                for fi39 := 0; fi39 < 19; fi39++ {
                    v4 -= -828
                    _ = fi39
                }
                _ = v38
            }
            _ = v37
        } else {
            v40 := []int{v12, (^(v12))}
            v16 = ("a" + v16)
            if (905 >= 42) {
                fmt.Println(v12, v16, v10)
            }
            _ = v40
        }
        var v41 float32 = 1e10
        _ = v41
    }
    v4 = -309
    var v42 int = v12
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v42
}
