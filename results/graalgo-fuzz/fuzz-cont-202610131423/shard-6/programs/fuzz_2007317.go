package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS2 struct { F0 string; F1 int32; F2 int16 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint64; F1 string; F2 uint }
func genFn0() uint8 {
    v5 := []int{}
    for fi6 := 0; fi6 < 10; fi6++ {
        v5 = append(v5, (-(fi6)))
        fmt.Println(fi6)
        _ = fi6
    }
    var v7 string = "hello"
    if false {
        switch 1 {
        case 32768:
            var v8 string = v7
            _ = v8
        }
    } else {
        var v9 bool = true
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 int64 = 42
            v7 = (v7 + ("foo bar" + "a"))
            _ = fi10
            _ = v11
        }
        var v12 float64 = 1.5
        _ = v12
        _ = v9
    }
    fmt.Println(v7)
    _ = v5
    _ = v7
    return 255
}

func main() {
    var v1 int64 = 355
    v2 := []uint32{195, 0, 516, 2}
    v3 := []string{"hello", ("hello" + "a")}
    var v4 bool = true
    for fi13 := 0; fi13 < 16; fi13++ {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 int16 = 912
            var v16 float32 = -2.25
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    fmt.Println(v1)
    for fi17 := 0; fi17 < 11; fi17++ {
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 bool = (1 > 42)
            if (485 >= 255) {
                v4 = v4
                if ("hello" > "") {
                    var v20 int16 = 0
                    v21 := GenS1{F0: 42, F1: 100, F2: 255}
                    _ = v20
                    _ = v21
                }
            } else {
                v22 := GenS2{F0: "a", F1: 47, F2: 42}
                var v23 string = (("hello" + "hello") + (("foo bar" + "") + ("hello" + "foo bar")))
                if v19 {
                    fmt.Println(v1, v19, v1)
                    if false {
                        var v24 float32 = 0.1
                        fmt.Println(v19)
                        var v25 uint64 = 42
                        _ = v24
                        _ = v25
                    }
                    if true {
                        fmt.Println(v1, v4, v4)
                        var v26 uint64 = 4294967295
                        var v27 int8 = -1
                        v4 = true
                        _ = v26
                        _ = v27
                    }
                    v28 := GenS3{F0: v1, F1: 1, F2: 255}
                    _ = v28
                }
                _ = len(v2)
                var v29 float32 = 0.1
                _ = v22
                _ = v23
                _ = v29
            }
            var v30 int8 = 100
            var v31 int16 = 32767
            _ = fi18
            _ = v19
            _ = v30
            _ = v31
        }
        var v32 int64 = v1
        if len(v3) > 0 {
            v3[0] = ("foo bar" + "a")
        }
        _ = fi17
        _ = v32
    }
    var v33 float32 = (0.0 + 0.1)
    fmt.Println(v33, v1, v33)
    v1 += v1
    fmt.Println(v1)
    v33 += v33
    for fi34 := 0; fi34 < 17; fi34++ {
        switch fi34 {
        case -384:
            if len(v2) > 0 {
                v2[0] = 2
            }
        case 32768:
            v35 := GenS4{F0: -884, F1: 2}
            _ = v35
        case 4294967295:
            switch (fi34 ^ fi34) {
            case 0:
                if (0 >= 42) {
                    for fi36 := 0; fi36 < 19; fi36++ {
                        var v37 int64 = (-(v1))
                        var v38 uint8 = 0
                        _ = len(v2)
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                }
            default:
                v4 = true
            }
        }
        v39 := GenS5{F0: 4294967295, F1: "\n\t", F2: 321}
        _ = fi34
        _ = v39
    }
    if v4 {
        switch 0 {
        case -9223372036854775808:
            for fi40 := 0; fi40 < 19; fi40++ {
                v2 = append(v2, 42)
                for fi41 := 0; fi41 < 8; fi41++ {
                    v42 := GenS2{F0: ("hello" + (("hello" + "foo bar") + ("" + ("hello" + "foo bar")))), F1: -1, F2: -583}
                    v4 = (!v4)
                    _ = fi41
                    _ = v42
                }
                _ = fi40
            }
        case 695:
            v43 := GenS3{F0: (-(v1)), F1: 18446744073709551615, F2: 0}
            _ = v43
        case 128:
            fmt.Println(v33, v33)
        }
        v4 = (!(!v4))
    } else {
        switch -2147483648 {
        case 42:
            if len(v2) > 0 {
                v2[0] = 143
            }
        case 954:
            var v44 uint8 = 0
            _ = v44
        case 1:
            v33 -= 0.1
        }
    }
    v1 += (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
