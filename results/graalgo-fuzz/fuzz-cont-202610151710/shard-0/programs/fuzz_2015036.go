package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 int32 }
func genFn0(p0 int64) bool {
    fmt.Println(p0)
    p0 = p0
    var v5 string = "hello"
    for fi6 := 0; fi6 < 3; fi6++ {
        if true {
            if true {
                v7 := GenS1{F0: 449}
                p0 += -828
                _ = v7
            } else {
                v8 := []int{fi6, fi6, (-(fi6)), (fi6 * fi6)}
                _ = v8
            }
        }
        fmt.Println(fi6)
        switch (v5 + "foo bar") {
        case "foo bar":
            fmt.Println(v5)
        }
        _ = fi6
    }
    _ = p0
    _ = v5
    return (!(!((-2.25 < 0.1) || ((!(32767 != 569)) || true))))
}

func main() {
    v1 := []int{1, 490, -412, -649}
    var v2 string = "a"
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    switch -1 {
    case 1:
        delete(v3, v2)
    case 3:
        switch v2 {
        case "\n\t":
            v9 := GenS2{F0: 9223372036854775807, F1: 224}
            _ = v9
        case "foo bar":
            if (0 == 42) {
                switch 310 {
                case -9223372036854775808:
                    for fi10 := 0; fi10 < 5; fi10++ {
                        switch fi10 {
                        case 4294967295:
                            v2 = ((("foo bar" + v2) + v2) + v2)
                        case -9223372036854775808:
                            v2 = v2
                        case 42:
                            v2 = (("foo bar" + ("a" + v2)) + "a")
                        }
                        v11 := GenS3{F0: -2, F1: 0}
                        _ = fi10
                        _ = v11
                    }
                }
            }
        case "hello":
            switch 4294967295 {
            case -128:
                var v12 int = 1
                _ = v12
            case -1:
                fmt.Println(v2)
            }
        default:
            delete(v3, (v2 + v2))
        }
    default:
        for fi13 := 0; fi13 < 1; fi13++ {
            if (true || true) {
                var v14 int = fi13
                v14 += v14
                if ((1.5 / 1.0) != 1.5) {
                    var v15 int64 = 9223372036854775807
                    fmt.Println(v14, v15)
                    _ = v15
                }
                _ = v14
            } else {
                var v16 uint32 = 4294967295
                v2 = ("a" + v2)
                v17 := GenS4{F0: -32768, F1: 255}
                _ = v16
                _ = v17
            }
            _ = fi13
        }
    }
    fmt.Println(v2)
    v18 := GenS3{F0: 0, F1: -9223372036854775808}
    var v19 bool = false
    fmt.Println(v19, v19)
    switch 2147483647 {
    case 128:
        fmt.Println(v2, v19)
    default:
        for fi20 := 0; fi20 < 6; fi20++ {
            v2 = v2
            v2 = ("a" + (v2 + (v2 + "")))
            v19 = (false || ((!(v19 || ((((1.5 - 0.1) * 1.5) / 0.1) > float32((1e10 * float64((0.1 / 3.14159))))))) && true))
            _ = fi20
        }
    }
    if v19 {
        v21 := GenS1{F0: 65535}
        if false {
            var v22 uint8 = 2
            _ = v22
        } else {
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 int = -449
                _ = fi23
                _ = v24
            }
            switch -1 {
            case -470:
                for fi25 := 0; fi25 < 15; fi25++ {
                    var v26 int64 = 1
                    _ = fi25
                    _ = v26
                }
            case 128:
                v27 := GenS1{F0: 547}
                _ = v27
            case 1:
                v19 = (1e10 > (0.0 - 1e10))
            }
        }
        _ = v21
    }
    if true {
        fmt.Println(v19, v2)
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        v19 = (v19 || v19)
        v29 := GenS4{F0: 42, F1: 186}
        _ = fi28
        _ = v29
    }
    if (-32768 != -1) {
        v19 = (0.0 < (1e10 / (1e10 / 3.14159)))
        v2 = v2
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        v2 = "foo bar"
        v2 = v2
        _ = fi30
    }
    switch 0 {
    case 4294967295:
        delete(v4, -598)
    case 42:
        switch (v2 + v2) {
        case "a":
            v31 := v4[527]
            _ = v31
        case "":
            var v32 float64 = 1.5
            _ = v32
        default:
            v33 := GenS3{F0: 127, F1: -1}
            _ = v33
        }
    default:
        v19 = false
    }
    v34 := GenS5{F0: 734}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
