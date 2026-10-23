package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 65535
    var v3 int16 = -867
    var v4 uint64 = 1
    v5 := make(map[int]bool)
    fmt.Println(v3, v3)
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: v2, F1: 2, F2: 2147483647}
        v8 := GenS1{F0: v2, F1: 206, F2: 2147483648}
        delete(v1, "\n\t")
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: (v2 >> ((18446744073709551615) & 63)), F1: 32767, F2: 242}
    fmt.Println(v4, v2)
    v4 += (v4 >> (((v2 << ((1) & 63))) & 63))
    v10 := GenS1{F0: 18446744073709551615, F1: 32767, F2: 2147483647}
    switch ("hello" + ("" + "")) {
    case "a":
        if (1e10 > 1.5) {
            v11 := GenS1{F0: v2, F1: 970, F2: 4294967295}
            _ = v11
        } else {
            v12 := GenS1{F0: v2, F1: 32767, F2: 65535}
            v13 := []int{1, -49}
            _ = v12
            _ = v13
        }
    case "\n\t":
        for fi14 := 0; fi14 < 2; fi14++ {
            v2 = uint(v4)
            if false {
                v15 := []int{359, fi14, -9223372036854775808}
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 uint16 = 0
                    _ = fi16
                    _ = v17
                }
                v15 = append(v15, fi14)
                _ = v15
            } else {
                if false {
                    v2 = 841
                }
                var v18 uint = 2
                var v19 int8 = int8(v4)
                _ = v18
                _ = v19
            }
            fmt.Println(v2)
            _ = fi14
        }
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        switch fi20 {
        case 0:
            var v21 int64 = -1
            _ = v21
        case -9223372036854775808:
            v22 := []int{fi20, (fi20 - int(v3)), (fi20 / ((399) | 1))}
            _ = v22
        }
        var v23 uint32 = 42
        _ = fi20
        _ = v23
    }
    fmt.Println(v3, v4, v3)
    if false {
        var v24 bool = (!(128 < 29))
        _ = v24
    }
    v4 -= (v4 / ((2) | 1))
    var v25 string = ""
    v3 = v3
    if (v4 <= v4) {
        v26 := GenS1{F0: v2, F1: 255, F2: 65535}
        _ = v26
    }
    switch (v25 + (v25 + v25)) {
    case "hello":
        v2 -= v2
    case "foo bar":
        v3 = -974
    }
    var v27 uint = 0
    var v28 uint = (^(v27))
    v29 := GenS2{F0: 0, F1: 87, F2: 672}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
