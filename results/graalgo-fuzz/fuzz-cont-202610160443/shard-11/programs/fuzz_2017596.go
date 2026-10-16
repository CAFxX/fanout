package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint64, p1 int16, p2 string) uint8 {
    v6 := GenS1{F0: 4294967295}
    fmt.Println(p1)
    switch (p2 + (p2 + (((p2 + (p2 + "")) + "hello") + p2))) {
    case "\n\t":
        if false {
            if (-211 > -1) {
                for fi7 := 0; fi7 < 11; fi7++ {
                    p2 = "a"
                    _ = fi7
                }
                v8 := GenS1{F0: 65535}
                _ = v8
            }
        }
    case "":
        p1 = p1
    case "a":
        var v9 uint16 = 65535
        _ = v9
    default:
        var v10 int = 32767
        _ = v10
    }
    var v11 uint16 = 788
    p0 -= (p0 / ((4294967295) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    return 11
}
func genFn1() int64 {
    var v12 float32 = 3.14159
    if true {
        v12 = 3.14159
        var v13 int32 = 0
        _ = v13
    }
    _ = v12
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 127
    v3 := []string{"\n\t", "a", ("hello" + ("foo bar" + "")), ""}
    var v4 bool = (!true)
    v5 := []string{}
    if true {
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 uint64 = 107
            if true {
                for fi16 := 0; fi16 < 15; fi16++ {
                    fmt.Println(v2, fi14)
                    var v17 int64 = 42
                    var v18 int16 = -547
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            _ = fi14
            _ = v15
        }
    } else {
        fmt.Println(v4)
    }
    v19 := GenS1{F0: 208}
    v20 := v1[(("" + ((("foo bar" + "\n\t") + "foo bar") + "hello")) + "a")]
    var v21 int8 = 38
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v2, v20, v2)
    switch (v20 + ((v20 + "foo bar") + v20)) {
    case "\n\t":
        for fi22 := 0; fi22 < 2; fi22++ {
            v2 = (^(v2))
            var v23 uint16 = 227
            v2 = (v2 << ((42) & 63))
            _ = fi22
            _ = v23
        }
    }
    _ = len(v3)
    var v24 float32 = ((1e10 - float32((-2.25 * 3.14159))) + (0.0 - (1e10 - (0.0 - ((1.5 + ((float32(1.5) / 1.0) * 1.5)) * 3.14159)))))
    v21 = -128
    if ((v24 - (-2.25 + float32(float64((3.14159 + v24))))) >= (1e10 - (0.1 + v24))) {
        v20 = v20
        v25 := GenS1{F0: 2}
        v1["\n\t"] = "hello"
        _ = v25
    }
    var v26 float32 = (v24 / float32(-2.25))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
