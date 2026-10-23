package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 int; F2 int64 }
func genFn0() uint {
    for fi4 := 0; fi4 < 10; fi4++ {
        if false {
            var v5 bool = (3.14159 < (0.1 * 1.5))
            v5 = (!(!v5))
            _ = v5
        }
        _ = fi4
    }
    if true {
        var v6 int32 = 2147483647
        fmt.Println(v6)
        if (true && (true && (0 < 1))) {
            var v7 float32 = -2.25
            v8 := GenS1{F0: -2147483648, F1: (9223372036854775807 >= 9223372036854775807)}
            _ = v7
            _ = v8
        } else {
            v9 := GenS2{F0: 18446744073709551615}
            switch "hello" {
            case "hello":
                for fi10 := 0; fi10 < 8; fi10++ {
                    var v11 bool = (128 <= 134)
                    _ = fi10
                    _ = v11
                }
            default:
                v6 = (v6 | (v6 % ((v6) | 1)))
            }
            _ = v9
        }
        _ = v6
    } else {
        var v12 uint = 18446744073709551615
        fmt.Println(v12)
        switch 2147483647 {
        case 128:
            var v13 int64 = 1
            _ = v13
        case -362:
            fmt.Println(v12)
        case 127:
            fmt.Println(v12)
        }
        _ = v12
    }
    return 255
}
func genFn1(p0 uint8) int32 {
    var v14 int8 = -128
    switch -225 {
    case 32767:
        var v15 int = 9223372036854775807
        _ = v15
    case 4294967295:
        v16 := []int{-695, 1, 32768, -277}
        _ = v16
    }
    var v17 uint = 4294967295
    v17 = (v17 << ((65535) & 63))
    _ = p0
    _ = v14
    _ = v17
    return 853
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[string]string)
    var v3 bool = true
    v18 := GenS3{F0: "", F1: 42}
    if (!(79 < 100)) {
        if false {
            switch "hello" {
            case "a":
                fmt.Println(v3, v3)
            }
            if v3 {
                var v19 int = 9223372036854775807
                var v20 string = "foo bar"
                v3 = v3
                _ = v19
                _ = v20
            } else {
                for fi21 := 0; fi21 < 8; fi21++ {
                    var v22 bool = v3
                    _ = fi21
                    _ = v22
                }
                fmt.Println(v3, v3)
                v3 = v3
            }
        }
    } else {
        var v23 int32 = 0
        _ = v23
    }
    v3 = (!("foo bar" == ("hello" + ("a" + "\n\t"))))
    var v24 int8 = -42
    var v25 float64 = 1e10
    fmt.Println(v25)
    if (792 == 127) {
        v24 = v24
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        _ = len(v2)
        _ = fi26
    }
    v27 := GenS4{F0: v3}
    fmt.Println(v1, v25)
    if false {
        var v28 int8 = v24
        var v29 string = "hello"
        _ = v28
        _ = v29
    }
    var v30 uint64 = 18446744073709551615
    v24 = v24
    for fi31 := 0; fi31 < 12; fi31++ {
        v24 -= (v24 + (v24 + v24))
        if v3 {
            var v32 bool = ((^(fi31)) > int(v1))
            _ = v32
        } else {
            if v3 {
                v3 = (32767 == -143)
                fmt.Println(v30)
                v33 := []int{(fi31 * 32767), (fi31 | 9223372036854775807), (fi31 & fi31)}
                _ = v33
            } else {
                var v34 int32 = 42
                _ = v34
            }
            for fi35 := 0; fi35 < 2; fi35++ {
                if v3 {
                    delete(v2, "hello")
                    v36 := GenS1{F0: fi35, F1: v3}
                    _ = v36
                }
                v37 := GenS5{F0: (42 < 2147483648), F1: fi31, F2: 0}
                _ = fi35
                _ = v37
            }
        }
        v1 -= 255
        _ = fi31
    }
    var v38 int8 = 100
    switch 288 {
    case 32767:
        v38 = (v24 * -59)
    case 32768:
        v1 = 42
    default:
        v39 := v2["a"]
        _ = v39
    }
    var v40 int8 = v24
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v30
    _ = v38
    _ = v40
}
