package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 string; F1 uint64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    var v4 float64 = 0.0
    var v5 int64 = 0
    var v6 int = 4294967295
    var v7 uint64 = 1
    if true {
        if true {
            for fi8 := 0; fi8 < 9; fi8++ {
                if (0.1 == v4) {
                    var v9 int = (fi8 ^ v6)
                    var v10 uint16 = 42
                    v11 := GenS1{F0: 42, F1: 100}
                    _ = v10
                    _ = v11
                    _ = v9
                } else {
                    var v12 int = fi8
                    v4 = v4
                    _ = v12
                }
                _ = fi8
            }
            var v13 float32 = 0.0
            switch (-(v6)) {
            case 32767:
                v14 := GenS1{F0: 127, F1: 100}
                _ = v14
            }
            _ = v13
        } else {
            for fi15 := 0; fi15 < 16; fi15++ {
                v16 := GenS1{F0: 181, F1: int8(v7)}
                for fi17 := 0; fi17 < 13; fi17++ {
                    var v18 int8 = -24
                    _ = fi17
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
        }
        if false {
            var v19 bool = false
            var v20 uint = 0
            for fi21 := 0; fi21 < 10; fi21++ {
                delete(v2, (v6 | (fi21 >> (((v20 % (((^(v20))) | 1))) & 63))))
                v22 := v1["foo bar"]
                _ = fi21
                _ = v22
            }
            _ = v19
            _ = v20
        }
    }
    if false {
        fmt.Println(v5, v5, v5)
        var v23 int64 = v5
        _ = v23
    }
    v24 := GenS1{F0: 255, F1: 127}
    switch "" {
    case "hello":
        switch ("\n\t" + ("a" + "a")) {
        case "hello":
            v25 := GenS2{F0: ("foo bar" + ("\n\t" + "foo bar")), F1: v7}
            _ = v25
        case "\n\t":
            var v26 bool = (0 == 358)
            _ = v26
        case "foo bar":
            var v27 bool = (4294967295 <= 4294967295)
            _ = v27
        }
    }
    var v28 uint = 42
    v29 := GenS2{F0: ("" + "hello"), F1: v7}
    v30 := GenS2{F0: ((("foo bar" + "a") + "a") + ""), F1: v7}
    for fi31 := 0; fi31 < 15; fi31++ {
        v32 := []int{0, v6, fi31, -9223372036854775808}
        _ = fi31
        _ = v32
    }
    v5 -= v5
    var v33 uint32 = 289
    v6 += (-(v6))
    v7 += v7
    if true {
        for fi34 := 0; fi34 < 12; fi34++ {
            for fi35 := 0; fi35 < 17; fi35++ {
                var v36 int32 = -7
                if (!false) {
                    v33 = 1
                    if true {
                        var v37 uint64 = v7
                        _ = v37
                    }
                    v38 := []int{(v6 ^ int(v33)), int(v28), v6, v6}
                    v7 = v7
                    _ = v38
                } else {
                    switch fi35 {
                    case 42:
                        v39 := GenS1{F0: 1, F1: 92}
                        _ = v39
                    case 309:
                        v36 -= v36
                    }
                    v6 = fi35
                    v40 := []int{(fi34 & (v6 / ((fi35) | 1)))}
                    _ = v40
                }
                var v41 uint16 = 2
                var v42 float32 = 0.0
                v43 := []int{2147483647, -90}
                _ = fi35
                _ = v36
                _ = v41
                _ = v42
                _ = v43
            }
            v44 := []int{int(v5), 0}
            var v45 float64 = 0.1
            _ = fi34
            _ = v44
            _ = v45
        }
        var v46 uint = v28
        if true {
            switch "hello" {
            case "a":
                fmt.Println(v46, v46, v7)
            case "hello":
                v2[(v6 % ((v6) | 1))] = ("" + (((("hello" + "hello") + ("" + "foo bar")) + "") + ("foo bar" + ("" + "foo bar"))))
            case "foo bar":
                for fi47 := 0; fi47 < 5; fi47++ {
                    for fi48 := 0; fi48 < 12; fi48++ {
                        var v49 uint = v28
                        _ = fi48
                        _ = v49
                    }
                    fmt.Println(v5, v33, v6)
                    v50 := GenS2{F0: ("hello" + ("hello" + ("foo bar" + ""))), F1: v7}
                    _ = fi47
                    _ = v50
                }
            }
        }
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
