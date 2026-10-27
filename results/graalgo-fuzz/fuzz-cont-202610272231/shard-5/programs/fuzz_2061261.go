package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string; F1 uint; F2 int32 }
type GenS7 struct { F0 int64; F1 int; F2 uint }
type GenS8 struct { F0 int32; F1 int8 }
func genFn0(p0 uint64, p1 uint16) string {
    switch 127 {
    case 32767:
        fmt.Println(p1)
    case -1:
        var v7 float32 = -2.25
        _ = v7
    }
    v8 := []int{2147483647, 455, -678, -1}
    _ = p0
    _ = p1
    _ = v8
    return "\n\t"
}
func genFn1() int8 {
    v9 := GenS1{F0: 255, F1: 1}
    var v10 int64 = 42
    _ = v10
    _ = v9
    return 57
}

func main() {
    var v1 uint = 1
    v2 := make(map[string]string)
    var v3 int8 = 42
    var v4 int8 = (^(v3))
    var v5 float64 = (-2.25 * 0.0)
    var v6 uint = (v1 << ((0) & 63))
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS2{F0: -9223372036854775808}
        _ = fi11
        _ = v12
    }
    if true {
        var v13 string = "a"
        fmt.Println(v5)
        _ = v13
    } else {
        fmt.Println(v5, v6)
        v14 := GenS3{F0: "foo bar", F1: 1}
        var v15 int = -1
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        if (float32(float64((float32(v5) / 1.5))) != 0.0) {
            var v17 uint64 = 18446744073709551615
            var v18 bool = ((32767 < 719) || (v5 >= v5))
            _ = v17
            _ = v18
        } else {
            var v19 int32 = 1
            for fi20 := 0; fi20 < 4; fi20++ {
                v21 := GenS1{F0: 127, F1: v6}
                v22 := GenS1{F0: 128, F1: v6}
                var v23 int64 = -493
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
            _ = v19
        }
        for fi24 := 0; fi24 < 4; fi24++ {
            if ((v3 ^ v3) <= genFn1()) {
                fmt.Println(v5, v4, fi24)
            }
            _ = fi24
        }
        delete(v2, ("foo bar" + "foo bar"))
        _ = fi16
    }
    v3 += v3
    if true {
        for fi25 := 0; fi25 < 6; fi25++ {
            for fi26 := 0; fi26 < 15; fi26++ {
                v4 -= -128
                _ = fi26
            }
            v27 := v2[""]
            _ = fi25
            _ = v27
        }
        v28 := GenS4{F0: 2}
        if (!(true && (true || ((true || (42 == 65535)) && true)))) {
            var v29 string = (("foo bar" + "\n\t") + "foo bar")
            v2[("hello" + v29)] = v29
            _ = v29
        } else {
            var v30 uint8 = 255
            _ = v30
        }
        _ = v28
    } else {
        var v31 int16 = 1000
        _ = v31
    }
    v32 := GenS5{F0: 42}
    v5 -= 0.0
    _ = len(v2)
    v5 -= (0.0 / 0.1)
    v33 := GenS6{F0: "foo bar", F1: (v1 & 2), F2: 938}
    var v34 uint16 = 32767
    if (-1 >= 34) {
        if false {
            v6 -= (^(v1))
            var v35 string = ""
            _ = v35
        }
    } else {
        var v36 int8 = 127
        v37 := GenS7{F0: -9223372036854775808, F1: 942, F2: v1}
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 9; fi38++ {
        v39 := GenS8{F0: 42, F1: genFn1()}
        fmt.Println(v5, v4)
        _ = fi38
        _ = v39
    }
    v3 = (v3 / ((int8(v6)) | 1))
    v6 += (^(v1))
    if false {
        v40 := []int{9223372036854775807, 32767, 824, -1}
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
