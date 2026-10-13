package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int }
func genFn0(p0 uint64, p1 string, p2 uint8) uint32 {
    v5 := []int{}
    var v6 int8 = -78
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 798
}

func main() {
    v1 := []string{"", "", "foo bar"}
    var v2 string = (("" + "a") + ("a" + ""))
    var v3 float64 = 0.0
    v4 := []uint32{}
    var v7 uint = 865
    for fi8 := 0; fi8 < 7; fi8++ {
        if false {
            v1 = append(v1, "a")
        }
        var v9 int32 = 0
        _ = fi8
        _ = v9
    }
    v3 += v3
    for fi10 := 0; fi10 < 4; fi10++ {
        v3 = (v3 * 1.5)
        _ = fi10
    }
    switch "a" {
    case "hello":
        v3 += v3
    default:
        for fi11 := 0; fi11 < 1; fi11++ {
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 uint64 = 18446744073709551615
                v2 = ((v2 + v2) + v2)
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 17; fi14++ {
                switch (fi11 * 128) {
                case -2147483648:
                    var v15 uint8 = 2
                    _ = v15
                }
                v3 = v3
                _ = fi14
            }
            _ = fi11
        }
    }
    _ = len(v4)
    var v16 uint32 = genFn0(4294967295, (v2 + "\n\t"), 188)
    v17 := GenS1{F0: 42, F1: "\n\t", F2: 154}
    for fi18 := 0; fi18 < 10; fi18++ {
        if (((3.14159 * float64(1e10)) - (0.1 / v3)) > (((v3 - (v3 - 0.1)) * -2.25) - v3)) {
            v3 -= v3
        }
        _ = fi18
    }
    var v19 uint16 = 683
    var v20 int32 = 1
    switch v2 {
    case "foo bar":
        fmt.Println(v7)
    default:
        switch (("foo bar" + v2) + ((((v2 + "foo bar") + (v2 + "a")) + v2) + ((v2 + v2) + "hello"))) {
        case "a":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            switch v2 {
            case "a":
                for fi21 := 0; fi21 < 8; fi21++ {
                    fmt.Println(fi21)
                    var v22 uint64 = 4294967295
                    _ = fi21
                    _ = v22
                }
            case "foo bar":
                v20 = v20
            }
        }
    }
    v19 += 12
    if true {
        if true {
            var v23 int32 = (-(v20))
            fmt.Println(v19, v23, v16)
            _ = v23
        }
        switch (v2 + (v2 + (v2 + v2))) {
        case "a":
            v24 := GenS2{F0: "hello"}
            _ = v24
        }
        for fi25 := 0; fi25 < 20; fi25++ {
            v19 -= 73
            v26 := GenS3{F0: (v7 & 42), F1: fi25}
            var v27 uint16 = v19
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    if false {
        v16 += (v16 + (v16 + v16))
    } else {
        _ = len(v4)
        var v28 uint32 = 4294967295
        var v29 int32 = v20
        _ = v28
        _ = v29
    }
    var v30 uint8 = 249
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 uint64 = 18446744073709551615
        v33 := GenS3{F0: (^(v7)), F1: fi31}
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v4
    _ = v7
}
