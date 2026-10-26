package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint8 }
func genFn0(p0 bool, p1 uint64, p2 uint16) uint16 {
    v4 := []int{127, 0}
    p0 = (-32768 < 1)
    if ((((-2.25 * (((-2.25 + 1.5) - 3.14159) - 1e10)) + 3.14159) <= 0.0) && false) {
        v4 = append(v4, 1)
        if ((97 == 1) && p0) {
            for fi5 := 0; fi5 < 7; fi5++ {
                fmt.Println(p2, p2, p2)
                var v6 uint = 42
                var v7 int16 = 42
                _ = fi5
                _ = v6
                _ = v7
            }
            switch ("hello" + "\n\t") {
            case "foo bar":
                _ = len(v4)
            }
            for fi8 := 0; fi8 < 20; fi8++ {
                for fi9 := 0; fi9 < 13; fi9++ {
                    p2 = p2
                    _ = fi9
                }
                _ = fi8
            }
        } else {
            v10 := GenS1{F0: -943}
            var v11 uint8 = 42
            _ = v10
            _ = v11
        }
    } else {
        var v12 uint64 = (p1 * 4294967295)
        if p0 {
            p0 = false
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        if ((3.14159 / ((float32((0.0 * ((((-2.25 / (0.1 * 1.5)) + float64(((0.0 / (-2.25 / 1.5)) * 1.5))) - -2.25) - ((((((1.5 / 0.1) * 0.0) * (3.14159 * 1e10)) / 1.0) / float64((0.1 + -2.25))) * (0.0 / -2.25))))) + 0.1) / 1e10)) < (-2.25 + (0.1 / 3.14159))) {
            var v13 uint32 = 278
            _ = v13
        } else {
            var v14 uint16 = 729
            _ = v14
        }
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}

func main() {
    var v1 int = -753
    var v2 string = "a"
    var v3 int = v1
    v2 = v2
    var v15 bool = false
    var v16 int32 = int32(v3)
    v17 := GenS1{F0: -1}
    fmt.Println(v2, v16)
    if v15 {
        if (!false) {
            if v15 {
                var v18 uint = 4294967295
                v19 := []int{(^(v1)), 42}
                if len(v19) > 0 {
                    v19[0] = -2147483648
                }
                _ = v18
                _ = v19
            }
            v20 := []int{(v1 - (v3 - 4294967295)), v1, v3, v3}
            var v21 string = "\n\t"
            _ = v20
            _ = v21
        } else {
            if (v15 && (255 <= 2)) {
                var v22 uint32 = 2147483647
                var v23 uint8 = 128
                _ = v22
                _ = v23
            }
            v24 := GenS1{F0: -1}
            _ = v24
        }
    } else {
        var v25 int8 = -30
        var v26 uint64 = 4294967295
        _ = v25
        _ = v26
    }
    var v27 int8 = 42
    for fi28 := 0; fi28 < 15; fi28++ {
        if (v27 >= (-(v27))) {
            v15 = true
            var v29 bool = (((v27 & (^(v27))) > (^(v27))) && (-1 >= 32767))
            var v30 int8 = 42
            _ = v29
            _ = v30
        } else {
            v27 -= 0
            switch "foo bar" {
            case "\n\t":
                var v31 float32 = 0.1
                _ = v31
            case "":
                v32 := GenS1{F0: 596}
                _ = v32
            default:
                v33 := GenS1{F0: -9223372036854775808}
                _ = v33
            }
        }
        for fi34 := 0; fi34 < 9; fi34++ {
            v35 := GenS2{F0: (v1 > 42)}
            _ = fi34
            _ = v35
        }
        _ = fi28
    }
    switch (v3 ^ 525) {
    case 226:
        var v36 bool = (!(225 > 2))
        _ = v36
    case 692:
        switch v2 {
        case "hello":
            for fi37 := 0; fi37 < 13; fi37++ {
                if v15 {
                    var v38 uint32 = 2
                    v3 -= v1
                    _ = v38
                }
                v39 := GenS3{F0: -2147483648, F1: 42}
                _ = fi37
                _ = v39
            }
        }
    }
    for fi40 := 0; fi40 < 5; fi40++ {
        v41 := GenS1{F0: -1}
        _ = fi40
        _ = v41
    }
    var v42 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v42
}
