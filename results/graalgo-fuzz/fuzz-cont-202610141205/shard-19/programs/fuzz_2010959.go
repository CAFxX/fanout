package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int64; F1 uint32; F2 int64 }

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]string)
    v3 := []uint32{2147483647, 979}
    var v4 string = (("" + (("a" + ("hello" + ("" + "\n\t"))) + "\n\t")) + ("a" + ("" + (("\n\t" + "") + ((("foo bar" + ("a" + "\n\t")) + "foo bar") + "a")))))
    var v5 uint = 956
    var v6 uint16 = 984
    for fi7 := 0; fi7 < 6; fi7++ {
        if true {
            fmt.Println(v4)
            for fi8 := 0; fi8 < 1; fi8++ {
                if len(v3) > 0 {
                    v3[0] = 414
                }
                _ = fi8
            }
            v6 += v6
        } else {
            v6 -= 32767
            var v9 uint16 = 65535
            v1 = v1
            _ = v9
        }
        _ = fi7
    }
    var v10 uint64 = (v1 >> (((^(v5))) & 63))
    fmt.Println(v4, v10, v10)
    if false {
        v11 := GenS1{F0: 279, F1: 2147483647}
        v6 = (v6 * 65535)
        _ = v11
    } else {
        var v12 int = -2147483648
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 13; fi13++ {
            v14 := GenS1{F0: 258, F1: 42}
            var v15 int32 = 0
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 int8 = -128
        v2[-885] = v4
        _ = v16
    }
    if ("\n\t" > ("\n\t" + (v4 + v4))) {
        if true {
            if true {
                for fi17 := 0; fi17 < 17; fi17++ {
                    if false {
                        var v18 int8 = 42
                        _ = v18
                    }
                    v19 := GenS1{F0: -1, F1: 478}
                    _ = fi17
                    _ = v19
                }
            }
            var v20 uint64 = v10
            _ = v20
        }
        if (0 > 101) {
            v3 = append(v3, 2147483647)
            if (2 == (v5 - 0)) {
                if (true && false) {
                    var v21 uint16 = 32767
                    var v22 uint16 = (v6 % (((v21 | 847)) | 1))
                    var v23 string = ((v4 + v4) + (v4 + v4))
                    _ = v21
                    _ = v22
                    _ = v23
                }
            } else {
                var v24 uint16 = (v6 >> (((v5 * (v5 / (((v5 >> ((2) & 63))) | 1)))) & 63))
                _ = v24
            }
        }
        var v25 uint16 = uint16(v1)
        _ = v25
    }
    v4 = "hello"
    v4 = (v4 + (v4 + v4))
    if (128 >= 236) {
        delete(v2, -1)
        v26 := GenS1{F0: 42, F1: -1}
        _ = v26
    }
    var v27 bool = ((100 > 73) || (0.1 <= (1e10 / (-2.25 / (-2.25 - 1.5)))))
    for fi28 := 0; fi28 < 8; fi28++ {
        v4 = (v4 + (v4 + v4))
        if v27 {
            var v29 uint = (v5 + 1)
            var v30 int32 = 0
            _ = v29
            _ = v30
        } else {
            var v31 int64 = -9223372036854775808
            v32 := GenS1{F0: 2147483647, F1: -2147483648}
            _ = v31
            _ = v32
        }
        _ = fi28
    }
    fmt.Println(v27)
    switch v4 {
    case "a":
        switch ((v4 + "") + (v4 + v4)) {
        case "foo bar":
            v1 = v10
        }
    case "hello":
        if v27 {
            fmt.Println(v10)
            var v33 uint = (^(v5))
            switch v4 {
            case "\n\t":
                var v34 int64 = -9223372036854775808
                _ = v34
            case "hello":
                if (v27 && (707 > 65535)) {
                    v4 = "\n\t"
                }
            case "":
                v35 := GenS2{F0: -110, F1: 796, F2: -215}
                _ = v35
            }
            _ = v33
        }
    }
    var v36 int16 = -32768
    v36 += v36
    v6 -= (v6 ^ v6)
    if v27 {
        for fi37 := 0; fi37 < 10; fi37++ {
            v38 := GenS1{F0: 880, F1: 59}
            _ = fi37
            _ = v38
        }
    }
    v2[9223372036854775807] = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
