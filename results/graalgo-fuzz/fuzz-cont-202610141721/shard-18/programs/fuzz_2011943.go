package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int64 }
func genFn0(p0 uint16, p1 string, p2 bool) uint32 {
    var v7 string = p1
    for fi8 := 0; fi8 < 16; fi8++ {
        if (!(0.1 < 0.1)) {
            v7 = ("foo bar" + v7)
        } else {
            for fi9 := 0; fi9 < 14; fi9++ {
                v7 = p1
                v10 := GenS1{F0: "hello", F1: 42, F2: -9223372036854775808}
                var v11 string = ("a" + ((v7 + (p1 + "\n\t")) + "foo bar"))
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        p2 = ((float64(0.1) - (0.1 - (0.0 * float64(0.0)))) > ((3.14159 - float64(1e10)) / 1.0))
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}
func genFn1(p0 uint32) uint8 {
    var v12 int16 = -32768
    p0 = (p0 << ((508) & 63))
    if (!(401 < -32768)) {
        if true {
            switch 42 {
            case 432:
                v13 := []int{}
                _ = v13
            case -550:
                var v14 uint8 = 59
                _ = v14
            }
            v12 -= (v12 % ((v12) | 1))
            var v15 int8 = -1
            _ = v15
        } else {
            if true {
                fmt.Println(v12)
                var v16 string = ("hello" + ("\n\t" + (("" + ("" + "a")) + ("\n\t" + "foo bar"))))
                _ = v16
            } else {
                for fi17 := 0; fi17 < 8; fi17++ {
                    var v18 int64 = int64(p0)
                    fmt.Println(fi17, fi17)
                    _ = fi17
                    _ = v18
                }
                v19 := []int{127, -181, 9223372036854775807, -2147483648}
                _ = v19
            }
            for fi20 := 0; fi20 < 12; fi20++ {
                v12 = 1000
                _ = fi20
            }
            var v21 uint16 = 1
            _ = v21
        }
        var v22 float64 = 3.14159
        _ = v22
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        v24 := []int{fi23, (fi23 + (fi23 - (fi23 / (((fi23 >> ((1) & 63))) | 1)))), fi23, fi23}
        _ = len(v24)
        _ = fi23
        _ = v24
    }
    _ = p0
    _ = v12
    return uint8(p0)
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 2
    var v3 int = 680
    v4 := make(map[int]string)
    v5 := make(map[int]int)
    var v6 float64 = 1.5
    v6 += (0.1 / float64(-2.25))
    var v25 int16 = -1
    v2 = v2
    switch "hello" {
    case "":
        var v26 int = v3
        _ = v26
    }
    var v27 uint32 = 2147483647
    var v28 int8 = 54
    fmt.Println(v27, v28, v25)
    switch (v3 + v3) {
    case 0:
        v6 -= 0.1
    case 32768:
        var v29 uint = 162
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
