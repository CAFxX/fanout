package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 uint64; F1 int32 }
type GenS7 struct { F0 uint; F1 int8; F2 bool }
func genFn0(p0 int16, p1 uint8, p2 int32) int {
    var v7 float32 = 3.14159
    switch "" {
    case "foo bar":
        v8 := GenS1{F0: (!("" < ("a" + "hello"))), F1: 1}
        _ = v8
    case "hello":
        v9 := GenS2{F0: p1}
        _ = v9
    default:
        var v10 string = "foo bar"
        _ = v10
    }
    v7 = float32(1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    var v1 uint = 202
    var v2 float32 = (-2.25 - 1e10)
    var v3 string = (("foo bar" + "hello") + "\n\t")
    var v4 uint = (v1 << (((^(v1))) & 63))
    var v5 float64 = ((-2.25 * ((float64(v2) + 1e10) / -2.25)) / float64(v2))
    var v6 int16 = 156
    var v11 int64 = 0
    for fi12 := 0; fi12 < 14; fi12++ {
        switch (("a" + "a") + v3) {
        case "foo bar":
            v1 += v4
        }
        _ = fi12
    }
    var v13 uint8 = 124
    fmt.Println(v13)
    var v14 float32 = v2
    var v15 uint16 = 1
    var v16 bool = (!true)
    var v17 float32 = (float32(v5) * -2.25)
    var v18 bool = (false || true)
    switch (v3 + "hello") {
    case "\n\t":
        var v19 int16 = 32767
        _ = v19
    case "hello":
        v20 := GenS3{F0: (uint8(v4) > (v13 % ((v13) | 1))), F1: ("a" + (("foo bar" + v3) + v3))}
        _ = v20
    case "":
        if (uint32(v4) <= 65535) {
            if true {
                var v21 uint16 = v15
                _ = v21
            }
        } else {
            var v22 uint32 = 2147483647
            v23 := GenS4{F0: (v22 >> ((4294967295) & 63))}
            _ = v22
            _ = v23
        }
    }
    v24 := GenS2{F0: (v13 & (v13 & 143))}
    v25 := GenS2{F0: uint8(v1)}
    switch v3 {
    case "":
        v26 := GenS5{F0: 32768, F1: 4294967295}
        _ = v26
    }
    if false {
        if v16 {
            if v16 {
                if ((v6 + 1) <= (v6 / ((32767) | 1))) {
                    var v27 int = 0
                    for fi28 := 0; fi28 < 5; fi28++ {
                        v29 := []int{1, (fi28 ^ 938), 862, (fi28 << ((42) & 63))}
                        switch ("\n\t" + (v3 + (((((v3 + "\n\t") + v3) + ((v3 + v3) + "a")) + ("" + "a")) + (v3 + ("" + v3))))) {
                        case "hello":
                            v27 = int(v11)
                        case "":
                            v1 += 42
                        }
                        _ = fi28
                        _ = v29
                    }
                    fmt.Println(v5, v11)
                    _ = v27
                } else {
                    v1 -= v4
                }
                var v30 uint64 = 42
                _ = v30
            } else {
                var v31 bool = (!(!((v4 + (v1 << ((65535) & 63))) < (v1 ^ uint(v13)))))
                v32 := []int{32768}
                _ = v31
                _ = v32
            }
            for fi33 := 0; fi33 < 10; fi33++ {
                switch fi33 {
                case -702:
                    fmt.Println(v17, v18)
                case 1:
                    for fi34 := 0; fi34 < 6; fi34++ {
                        fmt.Println(v2)
                        v2 -= 3.14159
                        _ = fi34
                    }
                case 0:
                    var v35 float32 = (0.0 - 1e10)
                    _ = v35
                }
                _ = fi33
            }
        }
    } else {
        if ("" != v3) {
            for fi36 := 0; fi36 < 14; fi36++ {
                v4 = (v4 + 2)
                _ = fi36
            }
            var v37 int64 = v11
            _ = v37
        } else {
            if (0 >= -128) {
                var v38 string = "foo bar"
                switch "\n\t" {
                case "\n\t":
                    for fi39 := 0; fi39 < 12; fi39++ {
                        for fi40 := 0; fi40 < 19; fi40++ {
                            fmt.Println(v3, v38, v38)
                            v41 := GenS6{F0: 4294967295, F1: -2147483648}
                            _ = fi40
                            _ = v41
                        }
                        var v42 uint16 = (v15 + v15)
                        _ = fi39
                        _ = v42
                    }
                default:
                    for fi43 := 0; fi43 < 4; fi43++ {
                        v5 += v5
                        v44 := []int{fi43}
                        v45 := GenS7{F0: v1, F1: -38, F2: v16}
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                }
                v46 := []int{-924, -661, 32767}
                _ = v38
                _ = v46
            }
            for fi47 := 0; fi47 < 6; fi47++ {
                var v48 int16 = v6
                _ = fi47
                _ = v48
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
