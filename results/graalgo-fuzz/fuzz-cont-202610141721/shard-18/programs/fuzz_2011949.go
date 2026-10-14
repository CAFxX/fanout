package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint16 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    var v3 int16 = -440
    v4 := []uint32{2147483647, 768, 0, 2}
    v5 := make(map[string]string)
    delete(v5, "")
    for fi6 := 0; fi6 < 9; fi6++ {
        v3 += (^(v3))
        _ = fi6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            if (false && (true && (true || (true || true)))) {
                if ((v3 & v3) > v3) {
                    var v9 int = fi7
                    v3 += -32768
                    _ = v9
                }
                if (0 >= 2147483647) {
                    v4 = append(v4, 42)
                    var v10 int32 = 0
                    v3 -= v3
                    _ = v10
                } else {
                    var v11 int = (-(fi8))
                    _ = v11
                }
            }
            v3 -= v3
            _ = fi8
        }
        _ = fi7
    }
    switch 417 {
    case 32767:
        v3 = (v3 >> ((255) & 63))
    default:
        switch "foo bar" {
        case "\n\t":
            if (false && (true && (2147483647 == 697))) {
                v3 = v3
            }
        case "":
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 uint32 = 93
                for fi14 := 0; fi14 < 7; fi14++ {
                    v15 := GenS1{F0: 2}
                    if (1e10 < 1.5) {
                        v16 := GenS2{F0: (9223372036854775807 == 136), F1: 127, F2: 2147483648}
                        _ = v16
                    }
                    v17 := GenS2{F0: (!true), F1: 42, F2: v13}
                    _ = fi14
                    _ = v15
                    _ = v17
                }
                _ = fi12
                _ = v13
            }
        }
    }
    var v18 uint64 = 4294967295
    fmt.Println(v18, v3)
    v19 := v2[4294967295]
    v20 := GenS3{F0: (v18 ^ 4294967295), F1: 0, F2: 255}
    v19 = (^(v19))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v21 := GenS1{F0: 144}
    if ((true && (870 != 0)) && ((2147483647 >= 0) || false)) {
        for fi22 := 0; fi22 < 5; fi22++ {
            if (42 < 0) {
                switch (fi22 << ((516) & 63)) {
                case 4294967295:
                    fmt.Println(v18, fi22, v19)
                }
                var v23 bool = false
                if (!((-539 != 1) && (!(v23 && (42 == 100))))) {
                    v24 := GenS1{F0: 4294967295}
                    fmt.Println(v19)
                    v3 = 1
                    _ = v24
                } else {
                    _ = len(v4)
                    var v25 int32 = -1
                    v23 = v23
                    _ = v25
                }
                _ = v23
            }
            if (true && false) {
                for fi26 := 0; fi26 < 1; fi26++ {
                    v27 := GenS3{F0: 793, F1: 2, F2: 42}
                    _ = fi26
                    _ = v27
                }
            }
            v2[(fi22 * -701)] = (^(v19))
            _ = fi22
        }
        for fi28 := 0; fi28 < 11; fi28++ {
            if true {
                var v29 int8 = int8(v18)
                _ = v29
            } else {
                for fi30 := 0; fi30 < 7; fi30++ {
                    var v31 string = "foo bar"
                    _ = fi30
                    _ = v31
                }
                var v32 int16 = (-(v3))
                _ = v32
            }
            _ = fi28
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
