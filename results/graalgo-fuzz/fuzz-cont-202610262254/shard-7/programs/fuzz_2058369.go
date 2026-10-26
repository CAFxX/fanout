package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 uint8 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 string = "a"
    var v3 uint8 = 42
    var v4 uint8 = 128
    var v5 bool = (true && true)
    v6 := make(map[string]int)
    v7 := GenS1{F0: 1, F1: 42, F2: (-(v1))}
    if false {
        v4 += v3
        if (708 == 248) {
            v8 := v6[(v2 + "a")]
            _ = v8
        } else {
            v1 = v1
        }
        v9 := []int{0, 937, -9223372036854775808, 32767}
        _ = v9
    }
    if v5 {
        v10 := GenS2{F0: 594}
        var v11 string = "hello"
        _ = v10
        _ = v11
    }
    switch v2 {
    case "a":
        for fi12 := 0; fi12 < 9; fi12++ {
            if (899 == 18446744073709551615) {
                if v5 {
                    v13 := GenS3{F0: 505, F1: v4}
                    _ = v13
                }
            } else {
                if (!v5) {
                    v14 := GenS2{F0: 65535}
                    for fi15 := 0; fi15 < 6; fi15++ {
                        fmt.Println(fi12)
                        v16 := []int{(fi15 << ((103) & 63)), 32768, (fi12 << ((2) & 63))}
                        v17 := GenS2{F0: 0}
                        if true {
                            _ = len(v6)
                            v18 := GenS3{F0: 4294967295, F1: v4}
                            v2 = (((("" + ("foo bar" + v2)) + (v2 + (v2 + ""))) + "foo bar") + ("" + v2))
                            var v19 uint32 = 2147483647
                            _ = v18
                            _ = v19
                        }
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = v14
                }
                if ((float32(1.5) + float32(1e10)) >= 3.14159) {
                    switch (fi12 & 32767) {
                    case 857:
                        switch fi12 {
                        case 32767:
                            var v20 uint64 = 42
                            _ = v20
                        case 4294967295:
                            fmt.Println(v3, v5, v1)
                        case 32768:
                            v5 = (!(-2.25 != (1.5 / 0.1)))
                        }
                    }
                    v1 = -9223372036854775808
                }
            }
            _ = fi12
        }
    }
    v3 = (^(v3))
    _ = len(v6)
    if v5 {
        var v21 uint8 = (v3 >> ((0) & 63))
        for fi22 := 0; fi22 < 14; fi22++ {
            for fi23 := 0; fi23 < 19; fi23++ {
                if v5 {
                    v2 = (("" + "foo bar") + v2)
                    var v24 uint64 = 42
                    _ = v24
                }
                _ = fi23
            }
            _ = fi22
        }
        fmt.Println(v4, v1)
        _ = v21
    } else {
        v25 := GenS2{F0: 255}
        v2 = ""
        _ = v25
    }
    v26 := GenS2{F0: 811}
    v4 += v4
    fmt.Println(v3)
    var v27 uint8 = (v3 & v4)
    for fi28 := 0; fi28 < 8; fi28++ {
        v29 := []int{(-(fi28)), fi28}
        _ = fi28
        _ = v29
    }
    v30 := GenS2{F0: 42}
    for fi31 := 0; fi31 < 6; fi31++ {
        v3 = (^(v27))
        var v32 int64 = v1
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
