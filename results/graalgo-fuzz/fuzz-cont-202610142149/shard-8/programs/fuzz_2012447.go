package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 int; F1 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 978
    var v3 int32 = 2147483647
    v2 = int(v3)
    var v4 float64 = 1.5
    var v5 uint32 = 400
    switch (v2 >> ((65535) & 63)) {
    case 2147483647:
        switch ("hello" + (("" + "") + "a")) {
        case "a":
            for fi6 := 0; fi6 < 15; fi6++ {
                v7 := []int{int(v3), fi6}
                v8 := GenS1{F0: ("hello" + ((("foo bar" + "foo bar") + "a") + ("hello" + ("a" + "hello"))))}
                _ = fi6
                _ = v7
                _ = v8
            }
        case "hello":
            v9 := []int{v2}
            _ = v9
        }
    case -2147483648:
        v10 := GenS2{F0: v5, F1: (((42 <= 0) || false) && (!(!true))), F2: 95}
        _ = v10
    }
    var v11 uint64 = 18446744073709551615
    switch (v2 & 0) {
    case 0:
        if false {
            fmt.Println(v11, v11)
            if false {
                switch v2 {
                case -462:
                    switch 32767 {
                    case 32767:
                        var v12 uint32 = (v5 % ((v5) | 1))
                        _ = v12
                    case -259:
                        v11 += (v11 % ((v11) | 1))
                    }
                case 2147483647:
                    var v13 int = (^(v2))
                    _ = v13
                case -225:
                    var v14 int = v2
                    _ = v14
                default:
                    var v15 int32 = (v3 >> ((0) & 63))
                    _ = v15
                }
            }
        } else {
            if true {
                v3 += (v3 & (^(v3)))
            } else {
                if false {
                    delete(v1, (v5 ^ (v5 & (^(v5)))))
                    var v16 string = "\n\t"
                    v11 -= 18446744073709551615
                    _ = v16
                }
                v4 = 1e10
            }
            v17 := []int{v2, (v2 << ((uint(v5)) & 63)), 596}
            if ((!false) && false) {
                v18 := GenS1{F0: (("foo bar" + "a") + "\n\t")}
                _ = v18
            }
            _ = v17
        }
    case -204:
        v19 := GenS1{F0: "\n\t"}
        _ = v19
    case 699:
        v20 := GenS3{F0: 42, F1: -9223372036854775808, F2: 779}
        _ = v20
    }
    var v21 float64 = float64(1e10)
    fmt.Println(v21)
    if (1e10 < 0.0) {
        v3 -= 1
        fmt.Println(v4, v4, v11)
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 uint32 = (v5 & (^(v5)))
            fmt.Println(v3, v2, v4)
            var v24 int16 = -1
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        if true {
            if (100 <= -92) {
                for fi26 := 0; fi26 < 19; fi26++ {
                    v11 = v11
                    v21 += v21
                    _ = fi26
                }
            } else {
                v27 := v1[v5]
                v4 = -2.25
                _ = v27
            }
            fmt.Println(v2, fi25)
        }
        for fi28 := 0; fi28 < 10; fi28++ {
            fmt.Println(fi28)
            v29 := []int{fi28, int(v3)}
            v11 -= v11
            _ = fi28
            _ = v29
        }
        _ = fi25
    }
    if (v2 != -1) {
        v11 = v11
        for fi30 := 0; fi30 < 18; fi30++ {
            v31 := GenS4{F0: -2147483648, F1: 101}
            for fi32 := 0; fi32 < 11; fi32++ {
                v33 := GenS2{F0: (v5 & (v5 / (((v5 * (v5 >> ((1) & 63)))) | 1))), F1: (42 <= 807), F2: uint8(v5)}
                fmt.Println(v4)
                if (true && (((1e10 / 1.0) + (0.1 * 1e10)) != float32(0.1))) {
                    var v34 float64 = 1e10
                    _ = v34
                } else {
                    if true {
                        v2 += -776
                    } else {
                        var v35 uint = 42
                        v1[(^(v5))] = (!(255 < 255))
                        _ = v35
                    }
                    fmt.Println(fi30)
                }
                _ = fi32
                _ = v33
            }
            for fi36 := 0; fi36 < 18; fi36++ {
                switch (fi36 << ((4294967295) & 63)) {
                case -37:
                    var v37 uint16 = 1
                    _ = v37
                }
                if false {
                    _ = len(v1)
                    v21 = v4
                } else {
                    v1[v5] = (0 <= 800)
                }
                var v38 float32 = (3.14159 + 0.0)
                _ = fi36
                _ = v38
            }
            _ = fi30
            _ = v31
        }
    }
    v39 := GenS5{F0: v2, F1: 1}
    v11 = v11
    var v40 uint64 = (^(v11))
    fmt.Println(v21, v5)
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
