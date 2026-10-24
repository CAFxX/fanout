package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
func genFn0(p0 uint64) int64 {
    if false {
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := []int{-397, (fi6 % (((fi6 + (^(fi6)))) | 1)), fi6}
            _ = fi6
            _ = v7
        }
        fmt.Println(p0)
    } else {
        p0 += (^(p0))
    }
    v8 := []int{32768, 454, 32768, 599}
    var v9 int64 = 9223372036854775807
    _ = len(v8)
    fmt.Println(p0, p0)
    _ = p0
    _ = v8
    _ = v9
    return v9
}
func genFn1(p0 int8) uint8 {
    v10 := []int{127, -1, 127, -886}
    if true {
        p0 += (p0 & p0)
    } else {
        switch "foo bar" {
        case "a":
            p0 -= p0
        case "":
            p0 = -21
        case "hello":
            p0 += (p0 >> ((255) & 63))
        }
    }
    if false {
        v11 := GenS1{F0: 283}
        v10 = append(v10, 985)
        _ = v11
    } else {
        var v12 int8 = (^(p0))
        if (!(!true)) {
            v13 := GenS2{F0: 4294967295}
            _ = v13
        }
        _ = v12
    }
    p0 -= 42
    var v14 int16 = 32767
    _ = p0
    _ = v10
    _ = v14
    return 128
}

func main() {
    var v1 string = ""
    var v2 uint = 1
    var v3 bool = (true || (1e10 < 0.1))
    var v4 uint = 2
    var v5 uint8 = 218
    v15 := []int{-2147483648, 127, 4294967295, 923}
    fmt.Println(v2, v5)
    var v16 int64 = -1
    var v17 float64 = 0.0
    var v18 uint64 = 42
    v19 := GenS1{F0: 65535}
    v15 = append(v15, 32768)
    v15 = append(v15, 128)
    v1 = ((v1 + "hello") + (v1 + (v1 + "hello")))
    fmt.Println(v18, v1, v4)
    switch "hello" {
    case "":
        for fi20 := 0; fi20 < 6; fi20++ {
            v17 -= v17
            if false {
                var v21 int = 0
                if len(v15) > 0 {
                    v15[0] = (fi20 >> ((42) & 63))
                }
                _ = v21
            }
            v22 := GenS3{F0: fi20}
            _ = fi20
            _ = v22
        }
    case "foo bar":
        var v23 bool = (1000 == 0)
        _ = v23
    case "hello":
        v2 -= v4
    }
    var v24 int8 = -20
    v25 := GenS1{F0: 2}
    if ((65535 >= 65535) && true) {
        fmt.Println(v3, v18, v2)
        switch (("\n\t" + ((v1 + ("a" + "")) + v1)) + ("hello" + (v1 + ("hello" + ((v1 + "a") + (v1 + v1)))))) {
        case "":
            if v3 {
                if v3 {
                    var v26 int64 = (v16 >> ((v4) & 63))
                    if ((!v3) || v3) {
                        var v27 float64 = v17
                        _ = v27
                    } else {
                        fmt.Println(v2, v17, v1)
                    }
                    _ = v26
                } else {
                    v16 += (v16 + (v16 << ((18446744073709551615) & 63)))
                    var v28 string = (v1 + ((v1 + v1) + v1))
                    _ = v28
                }
                var v29 uint16 = 964
                _ = v29
            }
        }
    } else {
        for fi30 := 0; fi30 < 10; fi30++ {
            v3 = v3
            _ = fi30
        }
    }
    var v31 uint = v4
    for fi32 := 0; fi32 < 12; fi32++ {
        v33 := GenS3{F0: fi32}
        v34 := GenS3{F0: fi32}
        v5 = (v5 | (v5 * v5))
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 bool = false
    var v36 uint = uint(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
