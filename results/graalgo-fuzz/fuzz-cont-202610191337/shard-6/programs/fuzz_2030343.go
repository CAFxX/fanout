package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int32; F2 string }
func genFn0(p0 string) uint32 {
    if true {
        v5 := GenS1{F0: p0, F1: 0}
        switch 0 {
        case 2147483647:
            for fi6 := 0; fi6 < 6; fi6++ {
                for fi7 := 0; fi7 < 16; fi7++ {
                    v8 := GenS2{F0: "hello"}
                    p0 = ("hello" + (p0 + ("a" + ("" + (((p0 + "a") + ((p0 + (("a" + "foo bar") + ((p0 + p0) + p0))) + p0)) + "foo bar")))))
                    _ = fi7
                    _ = v8
                }
                fmt.Println(p0)
                p0 = ((p0 + p0) + (p0 + p0))
                _ = fi6
            }
        case 336:
            switch (("\n\t" + "a") + (p0 + p0)) {
            case "\n\t":
                v9 := []int{128, 9223372036854775807}
                _ = v9
            }
        case 128:
            for fi10 := 0; fi10 < 12; fi10++ {
                var v11 int16 = -32768
                switch 1 {
                case 337:
                    fmt.Println(v11, p0, v11)
                }
                if (337 > 2147483647) {
                    for fi12 := 0; fi12 < 6; fi12++ {
                        var v13 uint64 = 18446744073709551615
                        var v14 uint32 = 594
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                }
                _ = fi10
                _ = v11
            }
        default:
            fmt.Println(p0)
        }
        _ = v5
    } else {
        v15 := []int{}
        _ = v15
    }
    var v16 int = 622
    fmt.Println(v16, p0)
    switch (-(v16)) {
    case -2147483648:
        var v17 int16 = -55
        _ = v17
    }
    p0 = (p0 + ((p0 + p0) + "foo bar"))
    _ = p0
    _ = v16
    return 42
}

func main() {
    v1 := []string{"\n\t", "a", ""}
    var v2 int16 = 32767
    var v3 uint64 = 42
    v4 := make(map[int]int)
    var v18 int32 = -566
    for fi19 := 0; fi19 < 10; fi19++ {
        if len(v1) > 0 {
            v1[0] = ("a" + "a")
        }
        v20 := GenS3{F0: 255, F1: v18, F2: ""}
        v21 := v4[fi19]
        _ = fi19
        _ = v20
        _ = v21
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = v3
    switch "" {
    case "a":
        v1 = append(v1, ((("\n\t" + "\n\t") + "") + ("foo bar" + ("foo bar" + ""))))
    }
    v22 := GenS3{F0: 127, F1: v18, F2: ""}
    if (42 > 42) {
        switch 2147483647 {
        case 591:
            v18 = (v18 % (((-(v18))) | 1))
        case 1:
            if len(v1) > 0 {
                v1[0] = (("foo bar" + (("hello" + "\n\t") + "foo bar")) + "hello")
            }
        case 0:
            var v23 string = ("foo bar" + "hello")
            _ = v23
        }
        v3 = v3
    }
    v24 := GenS2{F0: "foo bar"}
    var v25 int8 = 100
    for fi26 := 0; fi26 < 11; fi26++ {
        for fi27 := 0; fi27 < 17; fi27++ {
            switch 32768 {
            case 0:
                v2 = 42
            }
            var v28 uint8 = 128
            var v29 int32 = v18
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = len(v4)
        if (-2.25 > float32((0.0 * float64(0.1)))) {
            var v30 int16 = 130
            _ = v30
        } else {
            var v31 int16 = v2
            v32 := v4[(fi26 % ((508) | 1))]
            _ = v31
            _ = v32
        }
        _ = fi26
    }
    fmt.Println(v3, v3)
    var v33 float64 = 0.0
    v34 := GenS2{F0: ("\n\t" + "foo bar")}
    if false {
        fmt.Println(v18)
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
