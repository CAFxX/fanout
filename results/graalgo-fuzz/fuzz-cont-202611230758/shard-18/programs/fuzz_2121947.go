package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 string }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 4294967295
    var v3 uint64 = (v2 >> ((18446744073709551615) & 63))
    v4 := make(map[uint32]int)
    var v5 int16 = -1
    var v6 uint64 = v2
    v7 := GenS1{F0: false}
    for fi8 := 0; fi8 < 5; fi8++ {
        v4[65535] = 9223372036854775807
        v9 := []int{fi8, 641}
        _ = fi8
        _ = v9
    }
    if false {
        var v10 int8 = -128
        _ = v10
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            v12 := []int{4294967295, -9223372036854775808}
            var v13 uint64 = (^(v2))
            _ = fi11
            _ = v12
            _ = v13
        }
        if true {
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 string = "\n\t"
                var v16 uint = 991
                if (true || (-451 == -294)) {
                    fmt.Println(v5)
                    v17 := GenS1{F0: (false || (true && true))}
                    v6 += (v2 - v3)
                    _ = v17
                }
                _ = fi14
                _ = v15
                _ = v16
            }
            switch "\n\t" {
            case "":
                v18 := GenS1{F0: (-2.25 < -2.25)}
                _ = v18
            default:
                v5 += (v5 << ((2) & 63))
            }
            var v19 uint16 = 32767
            _ = v19
        }
    }
    v5 = v5
    for fi20 := 0; fi20 < 20; fi20++ {
        for fi21 := 0; fi21 < 6; fi21++ {
            if true {
                v22 := GenS1{F0: false}
                v23 := GenS2{F0: 9223372036854775807, F1: ("hello" + ("a" + ("foo bar" + "")))}
                _ = v22
                _ = v23
            }
            switch "hello" {
            case "":
                var v24 int = (fi21 / (((fi21 & (fi21 * 42))) | 1))
                _ = v24
            default:
                if (fi20 < (fi20 & (fi20 >> ((255) & 63)))) {
                    fmt.Println(fi20, v2)
                }
            }
            _ = fi21
        }
        switch (fi20 / (((fi20 << ((42) & 63))) | 1)) {
        case 57:
            switch "foo bar" {
            case "a":
                var v25 bool = false
                _ = v25
            default:
                var v26 uint32 = 2
                _ = v26
            }
        }
        v27 := []int{fi20}
        _ = fi20
        _ = v27
    }
    v28 := []int{327, 2147483647, -9223372036854775808, -9223372036854775808}
    var v29 bool = (42 != 127)
    switch 733 {
    case 32767:
        var v30 int16 = (v5 % (((v5 - (v5 + -32768))) | 1))
        _ = v30
    case 813:
        for fi31 := 0; fi31 < 8; fi31++ {
            v32 := GenS1{F0: ((int64(fi31) != -9223372036854775808) || (!(v29 || ((!v29) || ((^(v2)) >= v2)))))}
            _ = fi31
            _ = v32
        }
    case 1:
        fmt.Println(v29)
    }
    v33 := GenS2{F0: 0, F1: ((("" + "a") + "") + "\n\t")}
    v3 = 2
    var v34 uint32 = 935
    v6 = (v2 % (((v6 % ((v3) | 1))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
