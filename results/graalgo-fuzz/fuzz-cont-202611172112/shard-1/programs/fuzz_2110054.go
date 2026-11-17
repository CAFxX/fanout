package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint32, p1 uint8, p2 int) uint32 {
    var v5 uint64 = 975
    v6 := GenS1{F0: 0}
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int64 = 42
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p0
}
func genFn1(p0 uint8, p1 int16) uint64 {
    var v9 uint16 = 255
    for fi10 := 0; fi10 < 13; fi10++ {
        fmt.Println(p0, p1)
        fmt.Println(fi10, p0, p0)
        p0 -= 2
        _ = fi10
    }
    fmt.Println(v9, v9, p0)
    _ = p0
    _ = p1
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]string)
    var v3 uint = 958
    v4 := []int{2147483647, 32767, -2147483648}
    _ = len(v4)
    v3 = (v3 - (v3 % (((v3 + 18446744073709551615)) | 1)))
    var v11 int16 = 32767
    for fi12 := 0; fi12 < 16; fi12++ {
        fmt.Println(v3, v3, fi12)
        _ = fi12
    }
    if false {
        for fi13 := 0; fi13 < 20; fi13++ {
            if false {
                switch fi13 {
                case 2147483647:
                    v11 -= (v11 + v11)
                }
            } else {
                if true {
                    v11 += v11
                }
                v11 += (v11 % (((v11 * (^(v11)))) | 1))
                switch (fi13 >> ((v3) & 63)) {
                case 4294967295:
                    v3 += 384
                case 128:
                    var v14 uint8 = 128
                    _ = v14
                }
            }
            _ = fi13
        }
    } else {
        switch -144 {
        case -1:
            v15 := GenS1{F0: -9223372036854775808}
            _ = v15
        }
        var v16 uint32 = 923
        v17 := GenS1{F0: int64(v16)}
        _ = v16
        _ = v17
    }
    if false {
        for fi18 := 0; fi18 < 15; fi18++ {
            for fi19 := 0; fi19 < 13; fi19++ {
                v20 := GenS1{F0: 0}
                v11 = v11
                _ = fi19
                _ = v20
            }
            v3 = v3
            var v21 float32 = 0.1
            _ = fi18
            _ = v21
        }
        var v22 float32 = float32(1e10)
        _ = v22
    } else {
        var v23 string = ""
        switch (v23 + v23) {
        case "hello":
            _ = len(v1)
        case "foo bar":
            v3 = v3
        }
        _ = v23
    }
    var v24 int = -2147483648
    if (true || (!(float32(float64(((1.5 + 0.0) - ((-2.25 * 3.14159) + 1e10)))) >= 3.14159))) {
        var v25 int8 = 42
        var v26 int8 = 100
        var v27 uint16 = 32767
        _ = v25
        _ = v26
        _ = v27
    }
    v24 += -843
    if (!(true && false)) {
        v11 += (v11 & v11)
    } else {
        v3 = (v3 % ((v3) | 1))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
