package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() int16 {
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint16 = 2
        v9 := []int{2147483647}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 20}
    var v11 uint16 = 255
    var v12 float64 = 0.0
    var v13 float32 = 1.5
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 78
}
func genFn1(p0 uint8, p1 int16, p2 int32) int8 {
    switch 42 {
    case 128:
        if true {
            fmt.Println(p0)
            if (uint(p0) < 601) {
                var v14 int16 = (p1 % ((p1) | 1))
                _ = v14
            }
        } else {
            v15 := []int{-826, -281, 4294967295}
            for fi16 := 0; fi16 < 2; fi16++ {
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 int32 = -395
                    _ = fi17
                    _ = v18
                }
                var v19 float64 = 3.14159
                if false {
                    var v20 int16 = int16(p0)
                    _ = v20
                }
                _ = fi16
                _ = v19
            }
            p0 -= 0
            _ = v15
        }
    case -477:
        for fi21 := 0; fi21 < 13; fi21++ {
            for fi22 := 0; fi22 < 2; fi22++ {
                switch (fi22 & fi21) {
                case -45:
                    p0 = 1
                case 32767:
                    var v23 string = "foo bar"
                    _ = v23
                }
                p1 -= (p1 + p1)
                if false {
                    v24 := []int{}
                    fmt.Println(p1, fi22, fi21)
                    _ = v24
                }
                _ = fi22
            }
            var v25 uint32 = 2147483647
            var v26 uint32 = (v25 * 1)
            _ = fi21
            _ = v25
            _ = v26
        }
    }
    var v27 bool = true
    v28 := []int{42}
    fmt.Println(p0, p2, v27)
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v28
    return -128
}

func main() {
    var v1 int = 127
    var v2 bool = false
    v3 := []string{"a", "hello"}
    v4 := []string{(("\n\t" + "hello") + "\n\t"), (((("hello" + "a") + ((("\n\t" + ("hello" + "")) + ("\n\t" + "")) + "hello")) + "") + ""), "foo bar", "\n\t"}
    var v5 int64 = -9223372036854775808
    var v6 int8 = 127
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (65535 > 0) {
        switch "" {
        case "hello":
            var v29 float32 = 3.14159
            _ = v29
        default:
            var v30 float64 = 1e10
            _ = v30
        }
        for fi31 := 0; fi31 < 1; fi31++ {
            var v32 uint16 = 795
            var v33 int16 = 1
            v2 = v2
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    if false {
        var v34 uint8 = 186
        if true {
            if (!true) {
                fmt.Println(v6)
                v35 := GenS1{F0: v34}
                _ = v35
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v2 = (1 == 0)
                var v36 uint32 = 1
                _ = v36
            }
        }
        _ = v34
    }
    var v37 uint = 558
    if true {
        v38 := GenS1{F0: 42}
        _ = v38
    }
    if len(v4) > 0 {
        v4[0] = ("hello" + "")
    }
    fmt.Println(v1, v6, v5)
    v1 -= v1
    v39 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
