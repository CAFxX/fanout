package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int64, p1 uint64, p2 int16) uint32 {
    if (3.14159 > 1e10) {
        v5 := []int{32767, -9223372036854775808, -465, 313}
        _ = v5
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := []int{}
        p2 = 1000
        if ((1.5 / 3.14159) < 1.5) {
            for fi8 := 0; fi8 < 16; fi8++ {
                var v9 float64 = float64(float32(1.5))
                _ = fi8
                _ = v9
            }
        }
        _ = fi6
        _ = v7
    }
    if false {
        for fi10 := 0; fi10 < 12; fi10++ {
            switch "foo bar" {
            case "":
                var v11 uint16 = 65535
                _ = v11
            case "foo bar":
                p1 = p1
            case "hello":
                var v12 string = "\n\t"
                _ = v12
            }
            _ = fi10
        }
    } else {
        p2 -= p2
        p2 = (p2 >> ((573) & 63))
        for fi13 := 0; fi13 < 15; fi13++ {
            switch ("foo bar" + ("foo bar" + "")) {
            case "":
                v14 := []int{(-(fi13)), 2147483647, (fi13 ^ -524)}
                _ = v14
            }
            _ = fi13
        }
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    v1 := []uint32{4294967295, 0}
    var v2 int = 42
    v3 := make(map[int]uint32)
    v4 := make(map[int]bool)
    for fi15 := 0; fi15 < 19; fi15++ {
        v2 += v2
        _ = fi15
    }
    if false {
        var v16 uint64 = 535
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 uint64 = (v16 + v16)
            var v19 float64 = 0.0
            _ = fi17
            _ = v18
            _ = v19
        }
        v20 := v4[267]
        _ = v16
        _ = v20
    }
    var v21 int8 = -115
    var v22 uint = 4294967295
    for fi23 := 0; fi23 < 17; fi23++ {
        v24 := GenS1{F0: "foo bar"}
        _ = fi23
        _ = v24
    }
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    if true {
        switch "a" {
        case "foo bar":
            var v25 int64 = 9223372036854775807
            _ = v25
        case "":
            switch ("foo bar" + ("\n\t" + "a")) {
            case "":
                _ = len(v1)
            }
        case "hello":
            var v26 float64 = (-2.25 / ((float64(1e10) + 0.0) * -2.25))
            _ = v26
        }
    } else {
        v27 := GenS1{F0: ""}
        _ = v27
    }
    v1 = append(v1, 466)
    var v28 string = "foo bar"
    var v29 int32 = 2147483647
    fmt.Println(v21, v21)
    for fi30 := 0; fi30 < 1; fi30++ {
        for fi31 := 0; fi31 < 20; fi31++ {
            switch (v28 + ("\n\t" + v28)) {
            case "\n\t":
                v4[(fi31 + fi30)] = false
            case "hello":
                var v32 int32 = v29
                _ = v32
            }
            v22 += v22
            _ = fi31
        }
        var v33 uint64 = 403
        _ = fi30
        _ = v33
    }
    for fi34 := 0; fi34 < 9; fi34++ {
        var v35 int16 = 42
        v28 = v28
        v36 := v4[v2]
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 int16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v37
    _ = v4
}
