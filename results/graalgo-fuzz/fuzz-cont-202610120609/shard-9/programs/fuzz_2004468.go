package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }

func main() {
    var v1 int = -923
    var v2 float64 = (-2.25 + -2.25)
    var v3 uint = 1
    var v4 int32 = 42
    var v5 uint16 = 214
    var v6 int32 = (v4 % (((v4 << ((v3) & 63))) | 1))
    for fi7 := 0; fi7 < 9; fi7++ {
        if (213 <= 42) {
            v3 = v3
            if (float64(0.1) < 0.0) {
                v4 = (v4 + v4)
            } else {
                v8 := []int{-915, fi7, int(v3)}
                _ = v8
            }
            fmt.Println(v2, v5, v5)
        } else {
            fmt.Println(v4, v5)
        }
        if true {
            var v9 uint = (v3 ^ v3)
            _ = v9
        }
        var v10 uint8 = 1
        _ = fi7
        _ = v10
    }
    if (true && (0 < -1)) {
        v11 := GenS1{F0: 1, F1: 18446744073709551615, F2: 42}
        var v12 int32 = v6
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 15; fi13++ {
            if false {
                switch "hello" {
                case "":
                    var v14 int = v1
                    _ = v14
                case "\n\t":
                    v15 := GenS1{F0: 255, F1: (^(v3)), F2: 185}
                    _ = v15
                case "hello":
                    var v16 uint64 = 1
                    _ = v16
                }
            } else {
                v17 := []int{fi13, fi13}
                _ = v17
            }
            _ = fi13
        }
        v18 := []int{(v1 << ((v3) & 63)), v1, (v1 + (v1 | 4294967295))}
        _ = v18
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 int = v1
        v4 -= v4
        v2 = v2
        _ = fi19
        _ = v20
    }
    v6 += 699
    switch ("a" + "hello") {
    case "foo bar":
        fmt.Println(v6, v6)
    default:
        if (1 != 0) {
            switch "foo bar" {
            case "hello":
                fmt.Println(v6, v4)
            default:
                if (702 >= 2) {
                    var v21 int8 = 127
                    _ = v21
                }
            }
            for fi22 := 0; fi22 < 4; fi22++ {
                var v23 float32 = ((1.5 + 0.0) * 1e10)
                _ = fi22
                _ = v23
            }
        }
    }
    v3 = v3
    v24 := GenS1{F0: 95, F1: (v3 ^ (v3 * (v3 * (v3 + (^(v3)))))), F2: 42}
    v5 = 42
    fmt.Println(v6, v2, v4)
    if false {
        v4 = v6
        switch "hello" {
        case "foo bar":
            if (255 == 255) {
                v5 = v5
                v25 := []int{(-(v1)), v1, (v1 + 624)}
                _ = v25
            }
        default:
            var v26 uint16 = (v5 - (v5 % ((v5) | 1)))
            _ = v26
        }
        v27 := []int{(v1 * (v1 & (v1 << ((v3) & 63)))), v1, v1, (v1 ^ -142)}
        _ = v27
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        v29 := GenS1{F0: 199, F1: v3, F2: 1}
        var v30 string = ("hello" + "hello")
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        for fi32 := 0; fi32 < 9; fi32++ {
            v33 := []int{864, 128, v1, (fi32 << ((v3) & 63))}
            _ = fi32
            _ = v33
        }
        if (!((3.14159 + 3.14159) == 3.14159)) {
            var v34 int64 = -9223372036854775808
            _ = v34
        }
        for fi35 := 0; fi35 < 14; fi35++ {
            if (42 == 4294967295) {
                v3 = (v3 - (v3 & v3))
            }
            var v36 int = (v1 ^ v1)
            v4 = v4
            _ = fi35
            _ = v36
        }
        _ = fi31
    }
    v3 = (v3 << ((uint(v5)) & 63))
    v4 -= v6
    fmt.Println(v6)
    fmt.Println(v2, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
