package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 bool, p1 int64, p2 uint64) uint16 {
    v7 := []int{}
    for fi8 := 0; fi8 < 3; fi8++ {
        switch "" {
        case "":
            var v9 int16 = -537
            _ = v9
        case "hello":
            p1 = p1
        default:
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := GenS1{F0: 251}
                switch ((("hello" + "hello") + "a") + "\n\t") {
                case "":
                    var v12 int16 = -32768
                    _ = v12
                case "a":
                    _ = len(v7)
                }
                _ = fi10
                _ = v11
            }
        }
        switch (fi8 % ((fi8) | 1)) {
        case -2147483648:
            if p0 {
                if (-2.25 == 0.0) {
                    if ((!true) || (10 == 42)) {
                        var v13 uint32 = 2147483647
                        fmt.Println(p0, p2, fi8)
                        _ = v13
                    } else {
                        p1 += p1
                    }
                }
                for fi14 := 0; fi14 < 6; fi14++ {
                    fmt.Println(p0)
                    p0 = (true || p0)
                    _ = fi14
                }
                v15 := GenS2{F0: "hello", F1: 0, F2: 255}
                _ = v15
            }
        case 2147483647:
            fmt.Println(fi8, p2, p0)
        default:
            v16 := GenS2{F0: "", F1: -316, F2: 32767}
            _ = v16
        }
        _ = fi8
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        for fi18 := 0; fi18 < 14; fi18++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi18
        }
        if (0 <= 501) {
            var v19 float64 = 0.1
            _ = v19
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi17
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        switch ("\n\t" + "\n\t") {
        case "":
            var v21 float32 = ((0.1 - (1e10 / -2.25)) * (0.1 - 0.0))
            _ = v21
        }
        _ = fi20
    }
    var v22 uint64 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v7
    return 2
}

func main() {
    var v1 float32 = ((-2.25 / (3.14159 / -2.25)) + -2.25)
    v2 := make(map[uint32]string)
    var v3 uint64 = 849
    var v4 int64 = 111
    var v5 bool = (float32((3.14159 + (float64(v1) - 0.0))) < v1)
    var v6 uint = 4294967295
    for fi23 := 0; fi23 < 4; fi23++ {
        v1 += v1
        _ = fi23
    }
    v1 = -2.25
    if v5 {
        var v24 uint64 = v3
        v5 = false
        _ = v24
    }
    if v5 {
        v1 -= v1
    }
    var v25 int16 = 270
    fmt.Println(v4, v25)
    if ((1 > (v4 - v4)) && (v5 && (false || v5))) {
        _ = len(v2)
    }
    switch 128 {
    case 428:
        v6 = (v6 << (((v6 | (v6 * 42))) & 63))
    default:
        v5 = (((true && v5) && (((273 > 376) || true) || ("" != "hello"))) || v5)
    }
    if (v4 < 0) {
        var v26 bool = v5
        v2[574] = "foo bar"
        _ = v26
    } else {
        var v27 int8 = 42
        _ = v27
    }
    v28 := GenS2{F0: ("a" + ("hello" + ((("a" + ("a" + "")) + ("a" + "\n\t")) + "foo bar"))), F1: (v25 | -1), F2: 555}
    v29 := GenS1{F0: 127}
    if v5 {
        v6 += (v6 + v6)
        delete(v2, 224)
    } else {
        v6 = 42
    }
    _ = len(v2)
    v30 := GenS3{F0: 0}
    v31 := GenS2{F0: "foo bar", F1: (^(v25)), F2: 853}
    var v32 int = 2147483647
    v6 += 1
    if true {
        switch "foo bar" {
        case "foo bar":
            switch v32 {
            case 32767:
                v33 := []int{127}
                _ = v33
            default:
                v34 := GenS3{F0: 152}
                _ = v34
            }
        }
        var v35 int = v32
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
