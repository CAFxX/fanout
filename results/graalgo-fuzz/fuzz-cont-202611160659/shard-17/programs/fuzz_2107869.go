package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 int64; F2 uint64 }
func genFn0(p0 uint8, p1 uint16) int {
    var v5 int16 = 1
    var v6 float64 = 1e10
    v6 = (v6 * float64(-2.25))
    if ((2 <= 42) || (!(32767 == v5))) {
        var v7 int8 = 127
        if ("hello" == "a") {
            v6 = (1.5 * 1.5)
        } else {
            for fi8 := 0; fi8 < 11; fi8++ {
                switch (fi8 - (fi8 * (fi8 % ((fi8) | 1)))) {
                case -2147483648:
                    v9 := GenS1{F0: -855}
                    _ = v9
                case -976:
                    switch ("foo bar" + "\n\t") {
                    case "a":
                        v10 := []int{(fi8 + -1), fi8, fi8}
                        _ = v10
                    }
                case 928:
                    p1 += (p1 >> ((255) & 63))
                }
                var v11 uint = 1
                v12 := GenS1{F0: 1}
                _ = fi8
                _ = v11
                _ = v12
            }
        }
        _ = v7
    } else {
        var v13 uint16 = p1
        _ = v13
    }
    var v14 int16 = (v5 << ((18446744073709551615) & 63))
    _ = p0
    _ = p1
    _ = v14
    _ = v5
    _ = v6
    return -9223372036854775808
}
func genFn1() uint32 {
    var v15 uint8 = 255
    if ((float64((1.5 * 0.1)) <= 1e10) || (!(!(991 >= 255)))) {
        var v16 int = -9223372036854775808
        v17 := GenS1{F0: genFn0(2, 737)}
        var v18 uint32 = 65535
        _ = v16
        _ = v17
        _ = v18
    } else {
        var v19 int64 = 1
        _ = v19
    }
    switch 8 {
    case 127:
        v20 := []int{}
        _ = v20
    case 2147483647:
        v21 := GenS1{F0: 97}
        _ = v21
    default:
        if (!false) {
            for fi22 := 0; fi22 < 5; fi22++ {
                fmt.Println(v15, fi22)
                v23 := GenS2{F0: 0, F1: 242, F2: v15}
                _ = fi22
                _ = v23
            }
        }
    }
    _ = v15
    return 466
}

func main() {
    var v1 uint64 = 0
    var v2 int8 = 100
    var v3 uint16 = 32767
    v4 := make(map[uint32]int)
    switch "" {
    case "foo bar":
        for fi24 := 0; fi24 < 9; fi24++ {
            var v25 float32 = 0.1
            v26 := GenS2{F0: -1, F1: 1, F2: 9}
            _ = fi24
            _ = v25
            _ = v26
        }
    case "a":
        switch ("\n\t" + "hello") {
        case "hello":
            switch (("\n\t" + "") + "") {
            case "a":
                v27 := []int{216, 32768}
                _ = v27
            }
        case "a":
            for fi28 := 0; fi28 < 18; fi28++ {
                var v29 uint8 = 127
                var v30 uint = 2
                _ = fi28
                _ = v29
                _ = v30
            }
        default:
            _ = len(v4)
        }
    default:
        for fi31 := 0; fi31 < 5; fi31++ {
            v2 = -90
            _ = fi31
        }
    }
    for fi32 := 0; fi32 < 4; fi32++ {
        if true {
            if (true || (0.0 <= (((-2.25 + (0.1 - 1.5)) / 1e10) + 1.5))) {
                v33 := GenS3{F0: 239}
                _ = v33
            }
            v3 += (^(v3))
        } else {
            v34 := GenS3{F0: 71}
            switch -2147483648 {
            case -9223372036854775808:
                var v35 uint = 2
                _ = v35
            }
            _ = v34
        }
        _ = fi32
    }
    if false {
        var v36 int32 = 1
        for fi37 := 0; fi37 < 4; fi37++ {
            var v38 int64 = -1
            _ = fi37
            _ = v38
        }
        _ = v36
    }
    fmt.Println(v1, v3)
    v39 := GenS2{F0: 9223372036854775807, F1: 2, F2: 127}
    fmt.Println(v3)
    v40 := GenS1{F0: 0}
    switch genFn0(127, (^(v3))) {
    case 127:
        var v41 uint32 = 42
        _ = v41
    default:
        switch (("" + "a") + "") {
        case "hello":
            var v42 int32 = 0
            _ = v42
        }
    }
    fmt.Println(v3)
    v1 -= (v1 / (((v1 + v1)) | 1))
    v43 := []int{int(v3), 32768, 42, 127}
    v44 := GenS2{F0: 9223372036854775807, F1: 2, F2: 0}
    var v45 int64 = 570
    if true {
        v2 -= (v2 | (v2 | (-(v2))))
    }
    v46 := GenS3{F0: 128}
    if (false && true) {
        v43 = append(v43, 2147483647)
        v47 := GenS4{F0: -32768, F1: (v45 * 42), F2: (v1 >> ((18446744073709551615) & 63))}
        _ = v47
    } else {
        v48 := GenS3{F0: 43}
        for fi49 := 0; fi49 < 2; fi49++ {
            var v50 float32 = 1.5
            _ = fi49
            _ = v50
        }
        _ = v48
    }
    for fi51 := 0; fi51 < 8; fi51++ {
        v2 = v2
        _ = fi51
    }
    v52 := GenS3{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v52
}
