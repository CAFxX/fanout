package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int; F2 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 uint8; F2 bool }
func genFn0(p0 int32, p1 uint16) int {
    if true {
        switch "foo bar" {
        case "hello":
            if ((!(22 >= -9223372036854775808)) && (!false)) {
                var v4 int = -2147483648
                p1 = 42
                if false {
                    v5 := GenS1{F0: 18446744073709551615}
                    _ = v5
                }
                _ = v4
            } else {
                if true {
                    p1 -= p1
                } else {
                    fmt.Println(p1)
                    var v6 float32 = 3.14159
                    _ = v6
                }
                var v7 float64 = 1.5
                v7 = -2.25
                _ = v7
            }
        case "":
            for fi8 := 0; fi8 < 13; fi8++ {
                v9 := GenS2{F0: p1, F1: (fi8 << ((0) & 63)), F2: 0}
                for fi10 := 0; fi10 < 12; fi10++ {
                    v11 := GenS2{F0: 65535, F1: -17, F2: 9223372036854775807}
                    var v12 int8 = 127
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                p1 = 0
                _ = fi8
                _ = v9
            }
        case "foo bar":
            p1 += p1
        }
        if (((p0 ^ -2147483648) <= 614) && false) {
            p0 -= (^(p0))
            fmt.Println(p1, p0)
        } else {
            var v13 bool = (!(false || (1e10 >= (-2.25 / float64((float32(((0.1 / 3.14159) - ((1.5 / 3.14159) + 1.5))) / (1.5 + (-2.25 + float32((-2.25 / 1.0))))))))))
            p1 += (^(p1))
            p1 = p1
            _ = v13
        }
        v14 := []int{}
        _ = v14
    }
    p1 = 2
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return -740
}
func genFn1(p0 int64, p1 int) uint {
    v15 := []int{(p1 / (((p1 + (p1 >> ((965) & 63)))) | 1)), (p1 | p1), p1, genFn0(0, 445)}
    switch "\n\t" {
    case "foo bar":
        p1 -= genFn0(-1, 65535)
    case "\n\t":
        p0 = 5
    case "":
        var v16 int8 = 44
        _ = v16
    default:
        v17 := GenS3{F0: p0}
        _ = v17
    }
    v18 := GenS1{F0: 83}
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v18
    return 65535
}

func main() {
    var v1 int64 = 0
    var v2 string = (("hello" + (("hello" + "hello") + "a")) + "\n\t")
    var v3 uint8 = 218
    var v19 uint16 = 612
    v20 := GenS1{F0: 696}
    var v21 string = v2
    v19 -= (v19 | (^(v19)))
    var v22 int = 42
    v23 := GenS4{F0: (509 < 18446744073709551615)}
    switch v22 {
    case -2147483648:
        v24 := []int{(v22 << ((255) & 63)), v22, -814, v22}
        _ = v24
    case 32768:
        fmt.Println(v1, v3)
    }
    fmt.Println(v2, v22, v22)
    switch (v22 / (((-(v22))) | 1)) {
    case 475:
        v25 := GenS1{F0: 2}
        _ = v25
    }
    v3 += (v3 | (v3 * uint8(v19)))
    v22 += v22
    if false {
        if (0 != 1) {
            v26 := []int{(-(v22))}
            fmt.Println(v1)
            v19 = 2
            _ = v26
        }
    } else {
        v21 = (((v2 + v21) + v21) + v2)
        var v27 uint8 = v3
        if (42 < 0) {
            v1 += v1
            switch (^(v22)) {
            case 9223372036854775807:
                if (3.14159 < -2.25) {
                    switch ((("foo bar" + v21) + ((v21 + v21) + ("a" + ((v21 + "\n\t") + "foo bar")))) + (("" + (v21 + "\n\t")) + v21)) {
                    case "\n\t":
                        var v28 float64 = -2.25
                        _ = v28
                    case "a":
                        v21 = ("\n\t" + (v21 + v2))
                    }
                }
            case -9223372036854775808:
                v19 = v19
            case 32768:
                v29 := GenS3{F0: v1}
                _ = v29
            default:
                fmt.Println(v21, v27, v1)
            }
        }
        _ = v27
    }
    v21 = (((v2 + ((v2 + v2) + "\n\t")) + v2) + v2)
    v2 = (("a" + v21) + (("" + (("hello" + (v21 + v2)) + "\n\t")) + ""))
    if false {
        switch (v22 ^ v22) {
        case -1:
            v30 := GenS5{F0: v22, F1: (v3 << ((2) & 63)), F2: (!(-5 != 100))}
            _ = v30
        }
        if true {
            var v31 uint64 = 42
            _ = v31
        } else {
            if (v2 >= v2) {
                v32 := []int{genFn0(2147483647, 128), (v22 - (v22 / ((9223372036854775807) | 1))), 32768}
                _ = v32
            }
            fmt.Println(v19, v19)
            v19 = 255
        }
        if (true || true) {
            if (0.1 >= (3.14159 * (-2.25 * (1e10 * -2.25)))) {
                v33 := []int{}
                fmt.Println(v2)
                _ = v33
            } else {
                switch (v22 >> ((uint(v19)) & 63)) {
                case -559:
                    var v34 uint16 = v19
                    _ = v34
                default:
                    v35 := GenS1{F0: 2}
                    _ = v35
                }
            }
            if (685 <= 340) {
                if true {
                    var v36 int16 = -398
                    v19 += (v19 >> ((42) & 63))
                    var v37 int = (v22 / ((1) | 1))
                    _ = v36
                    _ = v37
                }
                v38 := []int{42}
                v39 := GenS1{F0: genFn1(v1, (v22 % ((-861) | 1)))}
                _ = v38
                _ = v39
            }
        }
    }
    switch (^(v22)) {
    case 794:
        var v40 string = v21
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
