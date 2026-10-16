package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
func genFn0(p0 int32, p1 bool) uint {
    p1 = (1 == 1)
    if false {
        p0 = p0
    } else {
        v4 := GenS1{F0: 65535, F1: 1}
        fmt.Println(p1)
        _ = v4
    }
    v5 := []int{32767, -887}
    _ = len(v5)
    _ = p0
    _ = p1
    _ = v5
    return 873
}

func main() {
    var v1 string = "foo bar"
    v2 := []string{v1, v1}
    v3 := []uint32{1, 954, 85}
    v1 = v1
    var v6 uint32 = 2
    var v7 string = (v1 + ("foo bar" + v1))
    v8 := GenS1{F0: 255, F1: 0}
    if (38 > 0) {
        if len(v3) > 0 {
            v3[0] = (v6 * (^(v6)))
        }
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        if true {
            for fi10 := 0; fi10 < 7; fi10++ {
                if (((v1 + (("" + (v1 + ((v1 + v1) + v1))) + v1)) + (v1 + ((v7 + ("hello" + v1)) + v7))) <= (v7 + v1)) {
                    v7 = "a"
                    fmt.Println(v6, fi9)
                    var v11 int = 243
                    _ = v11
                } else {
                    _ = len(v2)
                    v7 = v1
                }
                fmt.Println(v1, v6, fi10)
                var v12 string = ("foo bar" + "a")
                _ = fi10
                _ = v12
            }
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 float64 = (1.5 * (float64(float32(0.1)) * (0.1 + 1.5)))
                var v15 bool = (v14 >= 0.0)
                _ = fi13
                _ = v14
                _ = v15
            }
            if (!false) {
                fmt.Println(v7, v6)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
        _ = fi9
    }
    v6 += v6
    var v16 uint16 = 0
    fmt.Println(v1, v16, v1)
    switch (v7 + ("a" + ((v1 + ((("\n\t" + ("a" + ((v1 + ("a" + "")) + v7))) + (v1 + "foo bar")) + ("" + v1))) + ((((v1 + v1) + (("hello" + v1) + "hello")) + ("a" + ((v7 + ("foo bar" + (v7 + (("foo bar" + ("hello" + ("foo bar" + v1))) + (v7 + ("a" + v7)))))) + "\n\t"))) + ("\n\t" + "\n\t"))))) {
    case "a":
        var v17 bool = (true && ((("\n\t" + (v1 + (v7 + ""))) + v1) > (v7 + "foo bar")))
        _ = v17
    case "":
        v18 := GenS1{F0: 849, F1: -1}
        _ = v18
    case "\n\t":
        for fi19 := 0; fi19 < 13; fi19++ {
            v7 = "\n\t"
            if len(v2) > 0 {
                v2[0] = (v1 + (v1 + ((v1 + "foo bar") + v7)))
            }
            _ = fi19
        }
    }
    v6 += v6
    if (!(57 < 128)) {
        fmt.Println(v6)
        var v20 uint16 = (v16 * (v16 ^ 42))
        _ = v20
    }
    v7 = ("" + v7)
    _ = len(v3)
    if (!false) {
        if len(v3) > 0 {
            v3[0] = (v6 % ((v6) | 1))
        }
        v7 = "\n\t"
        v6 = v6
    }
    if (646 < 42) {
        var v21 uint8 = 127
        var v22 uint64 = 0
        _ = v21
        _ = v22
    } else {
        if len(v2) > 0 {
            v2[0] = (v7 + (("\n\t" + "") + v1))
        }
    }
    var v23 int32 = 0
    switch v7 {
    case "\n\t":
        for fi24 := 0; fi24 < 7; fi24++ {
            for fi25 := 0; fi25 < 7; fi25++ {
                _ = len(v2)
                var v26 uint16 = v16
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi25
                _ = v26
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v27 := GenS1{F0: (v16 << ((314) & 63)), F1: 127}
            _ = fi24
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
