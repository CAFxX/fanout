package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 uint8; F2 bool }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    var v4 uint16 = 504
    var v5 int32 = 2147483647
    var v6 uint64 = 317
    for fi7 := 0; fi7 < 12; fi7++ {
        v6 = (v6 + v6)
        v8 := []int{fi7, fi7, (fi7 % ((9223372036854775807) | 1)), fi7}
        var v9 bool = (true || ("hello" != (("foo bar" + ("hello" + "foo bar")) + "")))
        _ = fi7
        _ = v8
        _ = v9
    }
    if false {
        switch 2147483647 {
        case 1:
            var v10 uint32 = 4294967295
            _ = v10
        }
        v4 = v4
        switch "hello" {
        case "a":
            v3["hello"] = 65535
        }
    } else {
        if (!(-1 > -9223372036854775808)) {
            v11 := []int{1}
            var v12 uint = 65535
            v13 := GenS1{F0: true, F1: 1}
            _ = v11
            _ = v12
            _ = v13
        } else {
            fmt.Println(v4, v6, v6)
            var v14 uint64 = v6
            if (304 > 42) {
                for fi15 := 0; fi15 < 20; fi15++ {
                    for fi16 := 0; fi16 < 5; fi16++ {
                        fmt.Println(v1)
                        _ = fi16
                    }
                    _ = fi15
                }
                v4 -= v4
            } else {
                switch -765 {
                case -2147483648:
                    var v17 uint64 = (v14 ^ (v14 >> ((18446744073709551615) & 63)))
                    _ = v17
                }
            }
            _ = v14
        }
        v18 := GenS1{F0: false, F1: 783}
        _ = v18
    }
    fmt.Println(v6)
    if false {
        fmt.Println(v4, v1, v1)
    } else {
        for fi19 := 0; fi19 < 6; fi19++ {
            fmt.Println(fi19, v1)
            if (false && true) {
                v6 -= (v6 / ((v6) | 1))
                fmt.Println(fi19, v1)
                for fi20 := 0; fi20 < 5; fi20++ {
                    v21 := GenS1{F0: false, F1: 920}
                    var v22 int16 = 968
                    fmt.Println(v22, v22)
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            } else {
                var v23 int64 = -612
                if true {
                    var v24 uint32 = 0
                    switch (fi19 % ((fi19) | 1)) {
                    case 42:
                        v1 += (v1 - v1)
                    }
                    _ = v24
                } else {
                    v1 = -2.25
                }
                _ = v23
            }
            var v25 int32 = v5
            _ = fi19
            _ = v25
        }
        var v26 uint16 = v4
        fmt.Println(v4)
        _ = v26
    }
    switch "" {
    case "foo bar":
        delete(v3, "foo bar")
    case "\n\t":
        var v27 string = "foo bar"
        _ = v27
    case "a":
        if true {
            switch 32768 {
            case 32767:
                var v28 uint64 = 640
                _ = v28
            case 8:
                v29 := GenS2{F0: 45}
                _ = v29
            case -1:
                fmt.Println(v1)
            }
        } else {
            delete(v2, 1)
        }
    }
    switch 127 {
    case -44:
        switch "\n\t" {
        case "":
            v6 = 42
        case "foo bar":
            v30 := GenS1{F0: false, F1: 42}
            _ = v30
        }
    }
    if (0.1 <= v1) {
        for fi31 := 0; fi31 < 4; fi31++ {
            var v32 uint64 = v6
            _ = fi31
            _ = v32
        }
    } else {
        var v33 float32 = v1
        _ = v33
    }
    var v34 uint16 = (v4 * v4)
    v35 := GenS1{F0: (-20 == 967), F1: -9223372036854775808}
    v36 := GenS3{F0: -21, F1: 42, F2: (9223372036854775807 == 42)}
    v34 += (v4 & (v34 / (((v34 ^ (v34 ^ v34))) | 1)))
    var v37 float64 = float64((1e10 - 0.1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
