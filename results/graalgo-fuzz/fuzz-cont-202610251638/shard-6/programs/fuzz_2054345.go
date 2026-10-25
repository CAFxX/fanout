package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int64 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]int)
    var v3 uint64 = 85
    v4 := make(map[string]uint32)
    v5 := make(map[int]string)
    v6 := make(map[int]uint32)
    var v7 uint = 4294967295
    var v8 int8 = 1
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(v8, v8)
        switch "a" {
        case "hello":
            v7 -= 42
        }
        var v10 int = 577
        _ = fi9
        _ = v10
    }
    fmt.Println(v1, v1, v8)
    if (!false) {
        v2["a"] = 32768
    }
    var v11 float32 = 0.0
    v1 = 127
    switch ((("\n\t" + "a") + ("\n\t" + ("" + ("\n\t" + ("foo bar" + ("hello" + "foo bar")))))) + "") {
    case "":
        var v12 uint32 = 2147483648
        _ = v12
    default:
        fmt.Println(v8, v1)
    }
    if (31 > 2) {
        v11 = v11
        if ((32768 < -2147483648) || true) {
            v1 = 0
        }
        if false {
            fmt.Println(v11, v1)
        } else {
            v3 = v3
            switch "\n\t" {
            case "hello":
                var v13 uint8 = (v1 & (v1 & (v1 >> ((v7) & 63))))
                _ = v13
            case "foo bar":
                v8 = (v8 >> ((1) & 63))
            case "\n\t":
                for fi14 := 0; fi14 < 16; fi14++ {
                    if true {
                        v15 := GenS1{F0: (v3 * 0), F1: 9223372036854775807, F2: 9223372036854775807}
                        if (!false) {
                            fmt.Println(v8, v11, v11)
                        } else {
                            v16 := v2[""]
                            if (!(1 <= 9223372036854775807)) {
                                v1 = 0
                                var v17 uint8 = v1
                                _ = v17
                            }
                            _ = v16
                        }
                        _ = v15
                    } else {
                        v7 += v7
                        v18 := GenS1{F0: (^(v3)), F1: -664, F2: 1}
                        _ = v18
                    }
                    switch "hello" {
                    case "a":
                        if true {
                            switch fi14 {
                            case 0:
                                var v19 float32 = 1e10
                                _ = v19
                            case -98:
                                var v20 uint = (v7 * 4294967295)
                                _ = v20
                            case -552:
                                var v21 bool = (2147483647 == 2147483647)
                                _ = v21
                            }
                        } else {
                            if ((v3 + 890) > (v3 >> (((v7 + uint(v1))) & 63))) {
                                delete(v4, "foo bar")
                                if (165 >= (v1 * (v1 << (((^(v7))) & 63)))) {
                                    v22 := []int{}
                                    var v23 float32 = float32(0.0)
                                    _ = v22
                                    _ = v23
                                }
                            }
                            fmt.Println(v1, fi14)
                        }
                    case "":
                        for fi24 := 0; fi24 < 18; fi24++ {
                            var v25 uint32 = 2147483648
                            v3 = 1
                            v26 := GenS2{F0: 42}
                            if false {
                                v8 += -111
                            }
                            v25 += 589
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    }
                    var v27 int8 = (v8 & (v8 ^ v8))
                    _ = fi14
                    _ = v27
                }
            }
            switch 0 {
            case -1:
                _ = len(v6)
            }
        }
    } else {
        if (uint16(v1) > uint16(v7)) {
            for fi28 := 0; fi28 < 1; fi28++ {
                fmt.Println(fi28, v1, v8)
                var v29 string = ("\n\t" + (("a" + "\n\t") + "a"))
                if ((1 < -299) && (true && false)) {
                    v30 := GenS2{F0: v8}
                    _ = v30
                } else {
                    _ = len(v6)
                    for fi31 := 0; fi31 < 5; fi31++ {
                        v32 := GenS2{F0: v8}
                        v33 := GenS1{F0: v3, F1: -587, F2: 42}
                        v34 := GenS2{F0: (^(v8))}
                        _ = fi31
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                }
                _ = fi28
                _ = v29
            }
        }
        v35 := []int{42, 662}
        for fi36 := 0; fi36 < 13; fi36++ {
            v37 := v6[fi36]
            _ = fi36
            _ = v37
        }
        _ = v35
    }
    v1 = (v1 % ((v1) | 1))
    switch 700 {
    case 455:
        for fi38 := 0; fi38 < 20; fi38++ {
            v39 := []int{fi38}
            _ = len(v39)
            var v40 uint32 = 2147483648
            _ = fi38
            _ = v39
            _ = v40
        }
    case 127:
        _ = len(v2)
    case -2147483648:
        v7 = (v7 + 907)
    }
    for fi41 := 0; fi41 < 1; fi41++ {
        var v42 int8 = v8
        _ = fi41
        _ = v42
    }
    v3 -= v3
    v11 = ((v11 + 3.14159) / 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
