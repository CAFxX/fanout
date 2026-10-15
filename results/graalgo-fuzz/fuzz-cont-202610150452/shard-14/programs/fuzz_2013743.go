package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 string }
func genFn0(p0 int16) int32 {
    var v5 uint8 = 42
    v6 := []int{-793, 2147483647, -908}
    var v7 float32 = -2.25
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 79
}
func genFn1(p0 string, p1 uint32, p2 uint) int {
    if true {
        p0 = (((("\n\t" + p0) + "\n\t") + p0) + p0)
    } else {
        if false {
            if (85 >= 0) {
                var v8 uint8 = 123
                switch (p0 + "a") {
                case "foo bar":
                    p2 = p2
                case "":
                    p0 = p0
                }
                _ = v8
            } else {
                if false {
                    p1 += uint32(p2)
                } else {
                    for fi9 := 0; fi9 < 5; fi9++ {
                        if false {
                            if false {
                                if true {
                                    var v10 uint64 = 231
                                    v11 := []int{fi9}
                                    p2 += p2
                                    _ = v10
                                    _ = v11
                                }
                            }
                        }
                        fmt.Println(p2, p1, p2)
                        var v12 float32 = (((((float32(3.14159) - float32((float64(0.1) - (1.5 + 3.14159)))) - 0.1) + 1.5) * -2.25) * (float32(1e10) + 1e10))
                        _ = fi9
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 9; fi13++ {
                        if (true || (42 != 65535)) {
                            for fi14 := 0; fi14 < 7; fi14++ {
                                v15 := GenS1{F0: (p2 - (p2 - 18446744073709551615)), F1: 1}
                                var v16 string = p0
                                v17 := []int{}
                                _ = fi14
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 3; fi18++ {
                                var v19 int = (fi13 & 32767)
                                _ = fi18
                                _ = v19
                            }
                            v20 := []int{(fi13 ^ fi13), fi13, (fi13 >> (((p2 ^ p2)) & 63)), (fi13 * (fi13 | 127))}
                            _ = v20
                        }
                        switch "\n\t" {
                        case "a":
                            var v21 float64 = 0.1
                            _ = v21
                        }
                        for fi22 := 0; fi22 < 3; fi22++ {
                            p2 += (p2 & 580)
                            var v23 bool = ((2147483648 != p1) || true)
                            _ = fi22
                            _ = v23
                        }
                        _ = fi13
                    }
                }
            }
            v24 := GenS2{F0: p0}
            _ = v24
        }
        if (534 > uint16(p1)) {
            switch ("hello" + "") {
            case "a":
                var v25 uint32 = 65535
                _ = v25
            case "foo bar":
                for fi26 := 0; fi26 < 13; fi26++ {
                    v27 := GenS1{F0: (p2 << ((315) & 63)), F1: 1000}
                    fmt.Println(p1, p0)
                    _ = fi26
                    _ = v27
                }
            default:
                fmt.Println(p2)
            }
            v28 := GenS2{F0: ("\n\t" + (((("" + (((("" + (((((((("foo bar" + "foo bar") + (p0 + ((p0 + ("foo bar" + ((("a" + p0) + p0) + (p0 + "")))) + (p0 + p0)))) + (p0 + p0)) + (p0 + p0)) + p0) + p0) + "foo bar") + p0)) + p0) + p0) + p0)) + "foo bar") + "foo bar") + "hello"))}
            _ = v28
        }
    }
    if false {
        var v29 uint16 = 868
        _ = v29
    }
    _ = p0
    _ = p1
    _ = p2
    return -399
}

func main() {
    var v1 uint16 = 2
    var v2 string = "foo bar"
    var v3 uint64 = 18446744073709551615
    var v4 uint8 = 255
    var v30 uint8 = v4
    for fi31 := 0; fi31 < 13; fi31++ {
        switch ((v2 + "") + "") {
        case "foo bar":
            for fi32 := 0; fi32 < 6; fi32++ {
                v33 := []int{}
                _ = fi32
                _ = v33
            }
        case "hello":
            v34 := []int{(fi31 & (fi31 - 2147483647)), (fi31 - fi31), genFn1("hello", 732, 905), (fi31 % ((fi31) | 1))}
            _ = v34
        }
        _ = fi31
    }
    var v35 uint16 = (v1 & (v1 + v1))
    if (v2 <= (v2 + ((v2 + v2) + ("" + (v2 + (("a" + v2) + "")))))) {
        for fi36 := 0; fi36 < 7; fi36++ {
            v4 += 159
            _ = fi36
        }
    } else {
        fmt.Println(v4, v2, v35)
        v3 -= v3
    }
    v30 = v30
    fmt.Println(v35)
    v30 = 128
    v37 := []int{975, genFn1(v2, uint32(v3), 756)}
    var v38 uint16 = 42
    for fi39 := 0; fi39 < 4; fi39++ {
        _ = len(v37)
        _ = fi39
    }
    var v40 string = "hello"
    if ((!false) && true) {
        var v41 string = v2
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v40
}
