package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0() int {
    for fi4 := 0; fi4 < 5; fi4++ {
        fmt.Println(fi4)
        var v5 uint16 = 42
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: ""}
    if (514 == 2) {
        v7 := []int{-9223372036854775808, 4294967295, 32768, 202}
        _ = v7
    } else {
        v8 := []int{9223372036854775807, -233, 32767}
        _ = v8
    }
    var v9 uint64 = 122
    v9 = 0
    _ = v6
    _ = v9
    return 127
}
func genFn1(p0 uint16, p1 int) bool {
    switch (p1 << ((1) & 63)) {
    case 4294967295:
        if false {
            p1 -= 9223372036854775807
            v10 := GenS1{F0: "hello"}
            _ = v10
        } else {
            p0 = 1
        }
    }
    var v11 uint8 = 247
    if false {
        for fi12 := 0; fi12 < 17; fi12++ {
            for fi13 := 0; fi13 < 14; fi13++ {
                v14 := GenS1{F0: ("\n\t" + "hello")}
                v15 := GenS2{F0: 234, F1: 100}
                _ = fi13
                _ = v14
                _ = v15
            }
            if (!true) {
                fmt.Println(p1)
                var v16 uint64 = 360
                v17 := []int{}
                _ = v16
                _ = v17
            } else {
                if (563 >= 387) {
                    if (277 <= 1) {
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(fi12, p1, v11)
                        }
                        v18 := []int{}
                        if true {
                            p0 -= (^(p0))
                        }
                        _ = v18
                    }
                    if true {
                        for fi19 := 0; fi19 < 18; fi19++ {
                            var v20 uint = 255
                            p0 = (p0 % ((p0) | 1))
                            _ = fi19
                            _ = v20
                        }
                    } else {
                        var v21 float64 = -2.25
                        _ = v21
                    }
                }
                fmt.Println(fi12, v11, p1)
                if (1 == 55) {
                    var v22 bool = true
                    _ = v22
                } else {
                    v23 := []int{128}
                    _ = v23
                }
            }
            switch ("hello" + ("" + "\n\t")) {
            case "":
                p0 -= (p0 + uint16(v11))
            }
            _ = fi12
        }
    }
    var v24 int16 = 32767
    _ = p0
    _ = p1
    _ = v11
    _ = v24
    return (-9223372036854775808 == -568)
}

func main() {
    var v1 string = "\n\t"
    var v2 uint64 = 18446744073709551615
    var v3 uint = 18446744073709551615
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 int64 = int64(fi25)
        fmt.Println(fi25)
        _ = fi25
        _ = v26
    }
    var v27 uint16 = 2
    fmt.Println(v3)
    var v28 bool = true
    var v29 int64 = 42
    if v28 {
        v29 += v29
    }
    var v30 string = ""
    fmt.Println(v27, v3, v30)
    for fi31 := 0; fi31 < 19; fi31++ {
        v27 -= v27
        _ = fi31
    }
    switch 42 {
    case 4294967295:
        v32 := []int{-990}
        _ = v32
    default:
        switch v30 {
        case "a":
            v33 := []int{2147483647, genFn0(), -257, -2147483648}
            _ = v33
        case "hello":
            fmt.Println(v28, v2, v30)
        }
    }
    v27 -= (v27 << ((v3) & 63))
    v34 := []int{42, 32768}
    var v35 float32 = ((3.14159 / (1.5 + 3.14159)) * 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
}
