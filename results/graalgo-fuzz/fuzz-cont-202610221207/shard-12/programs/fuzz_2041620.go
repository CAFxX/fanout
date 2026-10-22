package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int16 }
func genFn0() uint16 {
    v4 := GenS1{F0: ("a" + "\n\t")}
    var v5 uint16 = 0
    _ = v4
    _ = v5
    return (v5 | (^(v5)))
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    var v3 bool = (1 >= 2)
    fmt.Println(v3)
    v3 = (!v3)
    if (0 >= -26) {
        switch "" {
        case "foo bar":
            v6 := v2[888]
            _ = v6
        default:
            v3 = (!(v3 || v3))
        }
        v3 = ((!((128 < 0) && false)) || ("a" == "\n\t"))
        if (!v3) {
            v3 = false
            for fi7 := 0; fi7 < 1; fi7++ {
                for fi8 := 0; fi8 < 2; fi8++ {
                    var v9 int = -2147483648
                    v10 := GenS1{F0: "\n\t"}
                    v9 = (fi8 ^ (fi8 - fi8))
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = fi7
            }
            var v11 uint16 = 447
            _ = v11
        } else {
            v3 = false
        }
    }
    switch (("foo bar" + "") + ("hello" + "")) {
    case "\n\t":
        if (v3 && (true && v3)) {
            var v12 int = 9223372036854775807
            var v13 uint64 = 295
            if v3 {
                v12 -= (v12 % ((32767) | 1))
            }
            _ = v12
            _ = v13
        }
    }
    var v14 uint8 = 255
    switch "\n\t" {
    case "a":
        switch "foo bar" {
        case "hello":
            var v15 uint = 941
            _ = v15
        case "\n\t":
            v1[4294967295] = -183
        }
    case "":
        for fi16 := 0; fi16 < 12; fi16++ {
            if (!v3) {
                v17 := GenS2{F0: 0}
                _ = v17
            } else {
                var v18 uint8 = (v14 << ((923) & 63))
                fmt.Println(v14, v18)
                _ = v18
            }
            if (true || v3) {
                for fi19 := 0; fi19 < 11; fi19++ {
                    var v20 float64 = 3.14159
                    var v21 int32 = 927
                    var v22 uint = 2
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
            }
            switch fi16 {
            case 2147483647:
                var v23 int16 = -1
                _ = v23
            case 128:
                v24 := []int{fi16, 32767, fi16, fi16}
                _ = v24
            }
            _ = fi16
        }
    default:
        v14 -= 128
    }
    v25 := v2[1]
    var v26 uint64 = 1
    if (32767 <= 967) {
        for fi27 := 0; fi27 < 16; fi27++ {
            if v3 {
                fmt.Println(v3)
            }
            switch (fi27 & fi27) {
            case -730:
                v14 = 2
            }
            v28 := []int{(-(fi27))}
            _ = fi27
            _ = v28
        }
        v29 := GenS3{F0: -393, F1: 627}
        if (-1 == 0) {
            v14 = 42
            v3 = (-141 <= -667)
        }
        _ = v29
    } else {
        v30 := GenS1{F0: (v25 + (v25 + v25))}
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
