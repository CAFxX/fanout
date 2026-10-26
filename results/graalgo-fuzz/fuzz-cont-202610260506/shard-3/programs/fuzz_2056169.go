package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 bool; F2 int32 }
type GenS5 struct { F0 int32; F1 uint32; F2 bool }

func main() {
    var v1 uint16 = 42
    var v2 float32 = 1e10
    var v3 uint64 = 42
    var v4 uint64 = v3
    var v5 int32 = -1
    var v6 uint = 18446744073709551615
    v7 := GenS1{F0: 127}
    for fi8 := 0; fi8 < 5; fi8++ {
        v6 -= (v6 & (v6 >> ((v6) & 63)))
        for fi9 := 0; fi9 < 5; fi9++ {
            v6 -= 618
            _ = fi9
        }
        var v10 uint64 = (v3 % ((v4) | 1))
        _ = fi8
        _ = v10
    }
    var v11 bool = (42 != 0)
    v5 = -2147483648
    var v12 int32 = (v5 << ((v6) & 63))
    fmt.Println(v4)
    v11 = v11
    var v13 uint32 = 42
    switch "hello" {
    case "a":
        v14 := []int{}
        _ = v14
    case "":
        var v15 float32 = v2
        _ = v15
    case "hello":
        switch "foo bar" {
        case "":
            v13 = 0
        case "foo bar":
            if true {
                for fi16 := 0; fi16 < 18; fi16++ {
                    var v17 bool = (fi16 <= int(v3))
                    switch (-(fi16)) {
                    case -2147483648:
                        var v18 int = (fi16 - fi16)
                        _ = v18
                    default:
                        var v19 uint8 = 36
                        _ = v19
                    }
                    v20 := GenS2{F0: (fi16 + (fi16 / (((fi16 - (-(fi16)))) | 1)))}
                    _ = fi16
                    _ = v17
                    _ = v20
                }
                v1 = 643
                v11 = false
            }
        }
    default:
        v21 := []int{int(v3), 4294967295, -903}
        _ = v21
    }
    switch 333 {
    case 116:
        v4 += v3
    default:
        for fi22 := 0; fi22 < 11; fi22++ {
            v23 := GenS2{F0: (fi22 % ((fi22) | 1))}
            switch ("\n\t" + "") {
            case "hello":
                var v24 uint32 = (v13 << ((v6) & 63))
                _ = v24
            }
            v25 := GenS1{F0: -91}
            _ = fi22
            _ = v23
            _ = v25
        }
    }
    var v26 bool = v11
    if (!false) {
        var v27 uint32 = (v13 >> ((2) & 63))
        for fi28 := 0; fi28 < 2; fi28++ {
            if (!false) {
                var v29 uint = v6
                _ = v29
            } else {
                var v30 int64 = -9223372036854775808
                _ = v30
            }
            var v31 int = 9223372036854775807
            _ = fi28
            _ = v31
        }
        _ = v27
    }
    var v32 string = ""
    switch v32 {
    case "\n\t":
        v26 = false
    case "a":
        for fi33 := 0; fi33 < 17; fi33++ {
            v32 = (v32 + "\n\t")
            _ = fi33
        }
    default:
        v34 := []int{-9223372036854775808, 66, 127, 42}
        _ = v34
    }
    var v35 int64 = -9223372036854775808
    var v36 uint64 = v3
    v2 -= ((v2 / v2) + 0.1)
    if false {
        v37 := GenS3{F0: 127}
        v38 := GenS4{F0: (v5 / ((v12) | 1)), F1: false, F2: 0}
        _ = v37
        _ = v38
    } else {
        switch v32 {
        case "a":
            v39 := []int{-343}
            _ = v39
        default:
            if (v4 <= (v36 % (((v4 / (((v3 + (v4 << ((v6) & 63)))) | 1))) | 1))) {
                v26 = v11
                switch (v32 + v32) {
                case "hello":
                    v36 += v4
                }
                v40 := GenS5{F0: v5, F1: v13, F2: v11}
                _ = v40
            } else {
                switch -9223372036854775808 {
                case 32768:
                    fmt.Println(v5)
                }
                v32 = ((v32 + (v32 + v32)) + v32)
            }
        }
        fmt.Println(v26, v3)
        v41 := []int{310, -474, -588, 128}
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
