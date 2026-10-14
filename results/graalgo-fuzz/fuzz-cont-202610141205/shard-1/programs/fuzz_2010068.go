package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint64; F1 string }
func genFn0() int16 {
    var v7 int = -2147483648
    v7 += (v7 + v7)
    for fi8 := 0; fi8 < 9; fi8++ {
        switch v7 {
        case -9223372036854775808:
            if (true || (100 != 100)) {
                if (1 < 4294967295) {
                    switch (v7 % (((v7 * (fi8 - 254))) | 1)) {
                    case 0:
                        var v9 int32 = -1
                        _ = v9
                    }
                    for fi10 := 0; fi10 < 11; fi10++ {
                        v7 -= (fi8 | fi8)
                        _ = fi10
                    }
                } else {
                    v11 := []int{(fi8 % ((v7) | 1)), v7}
                    var v12 uint64 = 42
                    _ = v11
                    _ = v12
                }
                v7 = (fi8 % ((-921) | 1))
                v13 := []int{fi8, v7}
                _ = v13
            }
        case -621:
            if false {
                var v14 uint16 = 1
                if true {
                    var v15 int8 = 42
                    _ = v15
                }
                if ("foo bar" <= "a") {
                    switch (("" + "") + ((("hello" + "foo bar") + "foo bar") + (("\n\t" + "") + "foo bar"))) {
                    case "foo bar":
                        v7 -= (fi8 ^ int(v14))
                    }
                    v16 := []int{}
                    _ = v16
                }
                _ = v14
            }
        }
        v17 := GenS1{F0: 352, F1: -665}
        _ = fi8
        _ = v17
    }
    _ = v7
    return 1000
}

func main() {
    var v1 int = 0
    var v2 int8 = 127
    v3 := []int{32768, 2147483647, v1, v1}
    var v4 int32 = 42
    v5 := make(map[string]uint32)
    var v6 int16 = 1000
    var v18 int = v1
    var v19 int32 = int32(v6)
    for fi20 := 0; fi20 < 10; fi20++ {
        for fi21 := 0; fi21 < 8; fi21++ {
            v18 += (v1 << ((1) & 63))
            _ = fi21
        }
        if true {
            for fi22 := 0; fi22 < 9; fi22++ {
                switch (-(fi20)) {
                case 655:
                    if true {
                        v19 = v4
                        v3 = append(v3, (v18 << ((255) & 63)))
                        fmt.Println(v6)
                    }
                case 2147483647:
                    fmt.Println(v4)
                default:
                    var v23 int32 = -2147483648
                    _ = len(v3)
                    v23 -= (v19 << ((682) & 63))
                    _ = v23
                }
                if true {
                    var v24 bool = ((v6 << ((4294967295) & 63)) != v6)
                    _ = v24
                } else {
                    var v25 uint = 1
                    _ = len(v5)
                    v19 += (v4 << (((v25 & (^(v25)))) & 63))
                    var v26 uint8 = 128
                    _ = v25
                    _ = v26
                }
                _ = fi22
            }
            fmt.Println(v1)
            for fi27 := 0; fi27 < 1; fi27++ {
                var v28 uint32 = 15
                v6 = (v6 | (v6 + (v6 & v6)))
                v29 := GenS1{F0: 470, F1: (v4 % (((v4 / ((2147483647) | 1))) | 1))}
                var v30 float64 = (1.5 + 1e10)
                var v31 int64 = 732
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
                _ = v31
            }
        }
        _ = fi20
    }
    v32 := GenS1{F0: 83, F1: int32(v2)}
    var v33 int32 = v4
    if (true && true) {
        var v34 int = (v18 & v18)
        if (!(65535 >= 255)) {
            switch "hello" {
            case "\n\t":
                var v35 uint32 = 0
                _ = v35
            default:
                var v36 uint = 107
                _ = v36
            }
            v37 := GenS2{F0: 961, F1: "hello"}
            if (!(true && false)) {
                if false {
                    _ = len(v5)
                }
            }
            _ = v37
        }
        v34 -= v34
        _ = v34
    }
    var v38 uint32 = 172
    if true {
        var v39 string = ""
        var v40 int = 4294967295
        _ = v39
        _ = v40
    }
    v38 = 2147483647
    v18 -= (v1 - -866)
    var v41 bool = ((v2 ^ 1) != (v2 * int8(v18)))
    v38 -= v38
    v19 += v4
    v6 += v6
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
