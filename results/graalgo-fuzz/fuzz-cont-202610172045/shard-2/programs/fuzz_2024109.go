package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() int32 {
    var v4 int = -2147483648
    fmt.Println(v4)
    _ = v4
    return -1
}

func main() {
    var v1 int16 = 0
    var v2 int16 = 42
    v3 := make(map[string]int)
    fmt.Println(v1, v2)
    if (-872 <= -995) {
        v1 = 617
        var v5 bool = (4294967295 == 2)
        _ = v5
    } else {
        v1 -= (v2 | v1)
        v1 += (v2 | (v2 >> ((956) & 63)))
        var v6 int = 127
        _ = v6
    }
    fmt.Println(v2)
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint8 = 2
        _ = fi7
        _ = v8
    }
    var v9 int32 = 163
    var v10 float64 = 3.14159
    fmt.Println(v2, v2, v1)
    var v11 int32 = v9
    for fi12 := 0; fi12 < 9; fi12++ {
        if ((59 <= -93) || false) {
            v9 = (^(v11))
        } else {
            switch fi12 {
            case 127:
                fmt.Println(v1, v1, fi12)
            }
            var v13 uint64 = 37
            for fi14 := 0; fi14 < 3; fi14++ {
                v13 -= 989
                _ = fi14
            }
            _ = v13
        }
        _ = fi12
    }
    v10 += 0.0
    switch -123 {
    case 629:
        var v15 bool = (!false)
        _ = v15
    case 232:
        v16 := []int{0}
        _ = v16
    case 451:
        _ = len(v3)
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        v18 := GenS1{F0: 4294967295}
        switch "hello" {
        case "a":
            v19 := []int{fi17, 4294967295, fi17, fi17}
            _ = v19
        case "":
            v20 := GenS1{F0: 18446744073709551615}
            _ = v20
        case "hello":
            v3[("foo bar" + "a")] = fi17
        }
        v21 := GenS1{F0: 255}
        _ = fi17
        _ = v18
        _ = v21
    }
    var v22 uint16 = 2
    var v23 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v9
}
