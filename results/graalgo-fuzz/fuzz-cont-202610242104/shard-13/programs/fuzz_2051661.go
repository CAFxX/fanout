package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint16; F1 int16 }
func genFn0(p0 int8) int8 {
    var v5 uint32 = 65535
    fmt.Println(p0)
    _ = p0
    _ = v5
    return (p0 | 100)
}
func genFn1() uint16 {
    var v6 uint = 255
    var v7 int = 496
    _ = v6
    _ = v7
    return 782
}

func main() {
    v1 := []string{("hello" + "\n\t"), "foo bar", "hello", "hello"}
    v2 := make(map[string]string)
    v3 := []string{"\n\t", "hello", ((("a" + "\n\t") + "a") + (("a" + "\n\t") + "foo bar")), "a"}
    v4 := make(map[uint32]int)
    var v8 bool = true
    if true {
        v1 = append(v1, ("a" + "foo bar"))
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 uint32 = 807
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 int16 = int16(fi11)
                fmt.Println(v12)
                v13 := GenS1{F0: (^(fi9)), F1: 0}
                _ = fi11
                _ = v12
                _ = v13
            }
            v3 = append(v3, "a")
            _ = fi9
            _ = v10
        }
    }
    v1 = append(v1, "foo bar")
    switch -2147483648 {
    case 9223372036854775807:
        _ = len(v1)
    }
    var v14 int32 = 2147483647
    var v15 string = ("hello" + (("foo bar" + ("\n\t" + "\n\t")) + (("hello" + ("hello" + "foo bar")) + "")))
    if true {
        delete(v4, 1)
    } else {
        switch 4294967295 {
        case 867:
            v15 = ((v15 + "foo bar") + v15)
        case -830:
            fmt.Println(v15)
        case 4294967295:
            v8 = (127 > -128)
        default:
            if true {
                var v16 float32 = -2.25
                for fi17 := 0; fi17 < 20; fi17++ {
                    v18 := GenS1{F0: (fi17 + (^(fi17))), F1: -9223372036854775808}
                    _ = fi17
                    _ = v18
                }
                v8 = (!(-332 <= 9223372036854775807))
                _ = v16
            }
        }
        if (!(68 <= -1)) {
            v19 := GenS2{F0: 255, F1: 0}
            _ = v19
        } else {
            v15 = ((v15 + ("" + v15)) + ("" + "foo bar"))
            var v20 uint16 = 255
            _ = v20
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if true {
        fmt.Println(v15)
        var v21 bool = ((2147483648 == 42) && true)
        _ = v21
    } else {
        var v22 float64 = (-2.25 * -2.25)
        for fi23 := 0; fi23 < 20; fi23++ {
            v3 = append(v3, v15)
            v14 = 0
            v14 -= v14
            _ = fi23
        }
        var v24 float64 = (v22 / (v22 + (v22 * v22)))
        _ = v22
        _ = v24
    }
    v8 = (0.1 >= ((1e10 * ((0.0 - 0.1) - 0.0)) / 1.0))
    var v25 string = "\n\t"
    var v26 uint8 = 127
    if true {
        var v27 uint = 359
        _ = v27
    }
    switch -517 {
    case 32767:
        for fi28 := 0; fi28 < 14; fi28++ {
            switch (fi28 << ((42) & 63)) {
            case 4294967295:
                if (!v8) {
                    var v29 float64 = -2.25
                    var v30 int16 = 32767
                    _ = v29
                    _ = v30
                }
            case 0:
                v4[420] = fi28
            }
            _ = fi28
        }
    default:
        var v31 int8 = -49
        _ = v31
    }
    if v8 {
        v32 := GenS1{F0: -529, F1: 121}
        v33 := GenS1{F0: 382, F1: -1}
        _ = v32
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v8
}
