package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 string, p1 string) uint8 {
    if true {
        var v7 bool = false
        var v8 float32 = 1.5
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 17; fi9++ {
            switch (fi9 * (^(fi9))) {
            case -732:
                switch 4294967295 {
                case -723:
                    v10 := GenS1{F0: fi9, F1: 0}
                    _ = v10
                case 127:
                    v11 := GenS2{F0: 1}
                    _ = v11
                }
            case 127:
                for fi12 := 0; fi12 < 13; fi12++ {
                    v13 := []int{32767, 128}
                    v13 = append(v13, fi9)
                    fmt.Println(fi9, fi12, p1)
                    _ = fi12
                    _ = v13
                }
            case 32767:
                v14 := []int{42, fi9, 1, 128}
                _ = v14
            }
            _ = fi9
        }
        if ((43 != 848) || (!false)) {
            p1 = "foo bar"
            if true {
                for fi15 := 0; fi15 < 11; fi15++ {
                    for fi16 := 0; fi16 < 6; fi16++ {
                        p1 = (p1 + "\n\t")
                        p0 = p1
                        var v17 int32 = 42
                        _ = fi16
                        _ = v17
                    }
                    fmt.Println(fi15, fi15)
                    _ = fi15
                }
            }
        } else {
            p1 = p0
        }
    }
    fmt.Println(p0, p1)
    var v18 uint = 2
    _ = p0
    _ = p1
    _ = v18
    return 128
}

func main() {
    var v1 int8 = 127
    v2 := []uint32{65535}
    var v3 uint32 = 900
    v4 := make(map[string]int)
    var v5 bool = false
    v6 := make(map[uint32]bool)
    for fi19 := 0; fi19 < 19; fi19++ {
        v3 -= (v3 & (^(v3)))
        v3 += v3
        _ = fi19
    }
    var v20 uint32 = v3
    v20 = v20
    var v21 uint = 65535
    v4["foo bar"] = 1
    fmt.Println(v1, v1, v21)
    if v5 {
        v20 = (^(v3))
        var v22 uint = 65535
        _ = v22
    }
    switch "hello" {
    case "\n\t":
        v23 := GenS1{F0: 16, F1: 9223372036854775807}
        _ = v23
    case "a":
        v24 := v4[("\n\t" + "")]
        _ = v24
    case "foo bar":
        fmt.Println(v20)
    }
    switch "" {
    case "a":
        for fi25 := 0; fi25 < 1; fi25++ {
            v20 = v3
            _ = fi25
        }
    case "\n\t":
        v1 += v1
    case "":
        var v26 uint64 = 0
        _ = v26
    }
    v20 = 794
    if false {
        if v5 {
            v27 := GenS1{F0: -9223372036854775808, F1: -9223372036854775808}
            _ = v27
        } else {
            for fi28 := 0; fi28 < 10; fi28++ {
                v20 -= (v3 % (((^(v20))) | 1))
                var v29 float32 = ((float32(0.1) / 1.0) * (float32(((0.1 / -2.25) / 1.5)) * 0.1))
                _ = fi28
                _ = v29
            }
            var v30 int32 = 414
            _ = v30
        }
        for fi31 := 0; fi31 < 2; fi31++ {
            v21 -= (^(v21))
            v1 = v1
            _ = len(v2)
            _ = fi31
        }
    }
    v2 = append(v2, v3)
    if (((((("a" + ("a" + "")) + "\n\t") + "foo bar") + "foo bar") + "hello") != "foo bar") {
        for fi32 := 0; fi32 < 3; fi32++ {
            var v33 uint8 = 1
            v20 = (v20 | (v20 * v20))
            v1 = v1
            _ = fi32
            _ = v33
        }
    }
    v3 -= (v3 << ((720) & 63))
    if (((!false) || (42 == -2147483648)) && v5) {
        switch -927 {
        case 31:
            if v5 {
                v34 := GenS2{F0: 940}
                _ = v34
            }
        case -905:
            for fi35 := 0; fi35 < 11; fi35++ {
                fmt.Println(v5)
                var v36 uint8 = 71
                _ = fi35
                _ = v36
            }
        }
        v3 = 818
    }
    v3 = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
