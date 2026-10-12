package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint8 }

func main() {
    var v1 int64 = -1
    var v2 uint32 = 268
    var v3 int32 = 555
    var v4 uint64 = 932
    switch "hello" {
    case "\n\t":
        fmt.Println(v3, v1, v4)
    case "a":
        for fi5 := 0; fi5 < 16; fi5++ {
            v6 := GenS1{F0: -481, F1: 339, F2: 48}
            for fi7 := 0; fi7 < 9; fi7++ {
                v8 := GenS1{F0: 0, F1: 1, F2: 42}
                _ = fi7
                _ = v8
            }
            v9 := []int{(fi5 ^ (fi5 ^ -2147483648)), (fi5 - (fi5 >> ((4294967295) & 63))), 1}
            _ = fi5
            _ = v6
            _ = v9
        }
    }
    v2 = uint32(v4)
    var v10 bool = false
    if (false && ((v1 + -553) >= 289)) {
        switch -572 {
        case 201:
            v11 := []int{42, 127}
            _ = v11
        case 920:
            var v12 int32 = (v3 - (v3 & (v3 & (v3 & (v3 - v3)))))
            _ = v12
        }
        var v13 uint32 = 2147483647
        var v14 uint16 = 105
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        if v10 {
            var v16 int64 = v1
            var v17 uint8 = 34
            _ = v16
            _ = v17
        }
        switch fi15 {
        case 2147483647:
            switch (fi15 / ((32767) | 1)) {
            case -1:
                v2 = (v2 >> ((4294967295) & 63))
            case 9223372036854775807:
                var v18 uint64 = v4
                _ = v18
            }
        default:
            v19 := []int{(fi15 | fi15), (fi15 >> ((uint(v4)) & 63)), (^(fi15)), fi15}
            _ = v19
        }
        v20 := []int{}
        _ = fi15
        _ = v20
    }
    if (3.14159 <= (((1.5 * float32(3.14159)) * ((0.1 + 0.0) + 1.5)) - 0.0)) {
        if ((725 >= 0) || true) {
            var v21 uint8 = 42
            for fi22 := 0; fi22 < 13; fi22++ {
                var v23 string = "hello"
                _ = fi22
                _ = v23
            }
            var v24 uint64 = (v4 * v4)
            _ = v21
            _ = v24
        } else {
            v25 := GenS1{F0: 0, F1: 50, F2: 215}
            if ((1 >= 18446744073709551615) || ((2 != 65535) && false)) {
                for fi26 := 0; fi26 < 12; fi26++ {
                    if v10 {
                        if v10 {
                            v27 := GenS1{F0: 1, F1: 44, F2: 108}
                            _ = v27
                        } else {
                            v28 := []int{}
                            _ = v28
                        }
                    } else {
                        v29 := []int{42, (fi26 >> ((86) & 63)), (fi26 + fi26), fi26}
                        _ = v29
                    }
                    _ = fi26
                }
            }
            v30 := GenS1{F0: -410, F1: 857, F2: 86}
            _ = v25
            _ = v30
        }
    }
    v31 := GenS1{F0: -1, F1: 4294967295, F2: 105}
    switch 903 {
    case -9223372036854775808:
        v32 := GenS1{F0: 32767, F1: 2, F2: 2}
        _ = v32
    default:
        v33 := []int{0, -9223372036854775808, -9223372036854775808}
        _ = v33
    }
    v4 = v4
    var v34 int = 0
    v34 = v34
    v3 = -1
    v35 := GenS1{F0: -29, F1: 1, F2: 174}
    v4 = v4
    for fi36 := 0; fi36 < 16; fi36++ {
        v37 := GenS1{F0: 135, F1: 18446744073709551615, F2: 0}
        var v38 int8 = -3
        for fi39 := 0; fi39 < 5; fi39++ {
            if (0 != -674) {
                var v40 uint16 = 605
                if (((!v10) && v10) && v10) {
                    v40 = v40
                }
                _ = v40
            } else {
                v10 = v10
            }
            v41 := []int{fi36, fi39, fi36}
            _ = fi39
            _ = v41
        }
        _ = fi36
        _ = v37
        _ = v38
    }
    var v42 int64 = v1
    v43 := GenS1{F0: 53, F1: 295, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v42
    _ = v43
}
