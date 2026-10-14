package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint, p1 int16) int16 {
    p0 = 1
    p1 = 1
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 int = -71
    var v2 uint16 = 248
    v3 := make(map[string]string)
    var v4 int64 = -1
    var v5 float64 = 0.1
    var v6 uint8 = 250
    if true {
        var v7 bool = (false && false)
        _ = v7
    } else {
        var v8 uint8 = (v6 ^ (^(v6)))
        var v9 string = "a"
        _ = v8
        _ = v9
    }
    v1 += v1
    var v10 float32 = ((0.1 - float32(((1.5 - v5) / v5))) / 1.0)
    switch "a" {
    case "\n\t":
        _ = len(v3)
    case "a":
        fmt.Println(v10, v10)
    }
    if false {
        var v11 uint = 42
        _ = v11
    } else {
        fmt.Println(v10)
    }
    v6 = v6
    for fi12 := 0; fi12 < 7; fi12++ {
        switch 1 {
        case -2147483648:
            v13 := GenS1{F0: 42, F1: 2}
            _ = v13
        }
        _ = fi12
    }
    v14 := []int{(v1 ^ -662), (v1 / ((v1) | 1)), v1, (v1 % (((v1 & v1)) | 1))}
    delete(v3, ("hello" + (("" + "a") + "foo bar")))
    for fi15 := 0; fi15 < 3; fi15++ {
        switch (-(fi15)) {
        case 42:
            var v16 uint = 660
            _ = v16
        case 32767:
            fmt.Println(v4)
        default:
            v1 += (v1 ^ (v1 << ((725) & 63)))
        }
        v6 = (^(v6))
        switch fi15 {
        case 971:
            for fi17 := 0; fi17 < 3; fi17++ {
                fmt.Println(fi15)
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                _ = fi17
            }
        case -300:
            _ = len(v14)
        case 4294967295:
            var v18 bool = ((("foo bar" + "foo bar") + ("a" + "hello")) < "hello")
            _ = v18
        default:
            switch ((("\n\t" + "hello") + (((("a" + ("\n\t" + ("\n\t" + "\n\t"))) + ((("" + "hello") + "foo bar") + "\n\t")) + (("a" + ("foo bar" + ("\n\t" + "a"))) + "foo bar")) + "hello")) + "") {
            case "\n\t":
                var v19 uint16 = (v2 - (v2 | (v2 << ((65535) & 63))))
                _ = v19
            case "a":
                if true {
                    var v20 int64 = v4
                    _ = v20
                }
            case "foo bar":
                v3["hello"] = (("\n\t" + "") + "\n\t")
            }
        }
        _ = fi15
    }
    fmt.Println(v4, v6)
    var v21 int8 = -1
    for fi22 := 0; fi22 < 14; fi22++ {
        if (618 == 1) {
            if (1000 != -1) {
                switch (v1 ^ (^(fi22))) {
                case 402:
                    switch (v1 / ((v1) | 1)) {
                    case -9223372036854775808:
                        fmt.Println(v1, v6)
                    case 4294967295:
                        var v23 int64 = 42
                        _ = v23
                    default:
                        var v24 int64 = (v4 - (-(v4)))
                        _ = v24
                    }
                }
            }
            fmt.Println(v2, v2)
        }
        var v25 int8 = v21
        _ = fi22
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
