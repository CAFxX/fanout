package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 string; F2 int64 }
func genFn0(p0 uint8, p1 int64, p2 int32) int32 {
    var v6 uint32 = 42
    v7 := []int{-449}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (p2 * p2)
}

func main() {
    v1 := []string{}
    var v2 uint16 = 827
    var v3 int = 144
    var v4 float64 = 3.14159
    v5 := make(map[string]bool)
    v4 += 1.5
    if true {
        v3 -= 32767
    } else {
        switch v3 {
        case -237:
            switch v3 {
            case -2147483648:
                if (true || (!true)) {
                    v3 = (v3 ^ 967)
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                    var v8 int16 = 42
                    _ = v8
                }
            case 78:
                v9 := GenS1{F0: 1}
                _ = v9
            case -438:
                var v10 uint64 = 0
                _ = v10
            }
        }
        var v11 int16 = 1000
        fmt.Println(v3)
        _ = v11
    }
    delete(v5, (((("foo bar" + ("foo bar" + "foo bar")) + "\n\t") + ("foo bar" + "a")) + ("" + "\n\t")))
    v2 = v2
    fmt.Println(v4, v3, v3)
    var v12 float64 = (0.1 * v4)
    for fi13 := 0; fi13 < 7; fi13++ {
        fmt.Println(v3)
        _ = fi13
    }
    var v14 uint8 = 42
    v3 -= -297
    for fi15 := 0; fi15 < 15; fi15++ {
        switch "hello" {
        case "foo bar":
            v16 := GenS2{F0: 9223372036854775807, F1: "", F2: 1}
            _ = v16
        case "a":
            delete(v5, ("" + (("a" + "hello") + ("hello" + ""))))
        }
        v12 = (3.14159 - float64(0.1))
        _ = fi15
    }
    var v17 float64 = (v12 + (v4 / float64((0.0 + float32(v4)))))
    switch (("" + ("foo bar" + ("\n\t" + "\n\t"))) + "") {
    case "a":
        for fi18 := 0; fi18 < 20; fi18++ {
            if (0.0 <= (((v12 / (((float64(((3.14159 - 1.5) - 3.14159)) * v17) + v17) * v17)) * (((float64(0.0) + float64(((0.1 + 0.1) - (0.1 - (-2.25 / (-2.25 / (1e10 / 0.1))))))) - ((0.0 - (v12 + float64(0.0))) / v4)) / 1.5)) * 3.14159)) {
                v14 += v14
                v3 += fi18
            } else {
                v5["foo bar"] = false
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                v4 = 0.1
            }
            v17 -= (float64(3.14159) / 1.0)
            switch (v3 << ((906) & 63)) {
            case 32768:
                _ = len(v1)
            case 1:
                if (!false) {
                    fmt.Println(fi18, v4, v17)
                    _ = len(v1)
                }
            case 128:
                var v19 uint64 = 0
                _ = v19
            default:
                fmt.Println(v3)
            }
            _ = fi18
        }
    case "hello":
        if (!((v14 <= 151) || true)) {
            v3 = v3
            _ = len(v5)
        } else {
            if false {
                v14 = (v14 & (v14 % ((v14) | 1)))
                fmt.Println(v17, v17, v2)
                for fi20 := 0; fi20 < 6; fi20++ {
                    for fi21 := 0; fi21 < 2; fi21++ {
                        v4 += v12
                        _ = fi21
                    }
                    _ = fi20
                }
            } else {
                v22 := GenS2{F0: 992, F1: "hello", F2: 9223372036854775807}
                if (v4 == (v4 / v12)) {
                    v17 = v12
                }
                var v23 int32 = -2147483648
                _ = v22
                _ = v23
            }
            var v24 string = "a"
            for fi25 := 0; fi25 < 4; fi25++ {
                v26 := GenS1{F0: 2147483647}
                v5[""] = (!(21 <= 100))
                fmt.Println(v2)
                _ = fi25
                _ = v26
            }
            _ = v24
        }
    default:
        if (42 >= v2) {
            _ = len(v5)
            v5["a"] = true
        } else {
            v3 -= (^(v3))
            var v27 bool = ((v12 / (0.1 / 1.5)) < (v4 * v17))
            _ = v27
        }
    }
    switch (v3 ^ v3) {
    case 42:
        for fi28 := 0; fi28 < 3; fi28++ {
            var v29 string = ("" + (("" + "\n\t") + "foo bar"))
            var v30 float64 = float64((0.1 - -2.25))
            v17 = v30
            _ = fi28
            _ = v29
            _ = v30
        }
    case 4294967295:
        for fi31 := 0; fi31 < 13; fi31++ {
            v32 := v5["a"]
            _ = fi31
            _ = v32
        }
    }
    v33 := v5[""]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
