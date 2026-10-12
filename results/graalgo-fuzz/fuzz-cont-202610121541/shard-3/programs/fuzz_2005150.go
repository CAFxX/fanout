package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int32, p1 uint16) uint8 {
    fmt.Println(p1, p1)
    p1 = 879
    fmt.Println(p1, p0)
    p0 -= (p0 / (((^(p0))) | 1))
    var v7 uint64 = 42
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1(p0 bool) string {
    if true {
        fmt.Println(p0)
        fmt.Println(p0)
    } else {
        if (4294967295 <= 0) {
            fmt.Println(p0)
            var v8 string = ((("foo bar" + "") + "hello") + "foo bar")
            _ = v8
        } else {
            v9 := []int{294}
            v10 := GenS1{F0: 0}
            p0 = (!(107 != 42))
            _ = v10
            _ = v9
        }
        v11 := []int{}
        v12 := GenS1{F0: 42}
        _ = v11
        _ = v12
    }
    var v13 uint8 = 128
    v14 := []int{-925}
    v13 = genFn0(42, 255)
    switch 89 {
    case 42:
        _ = len(v14)
    }
    _ = p0
    _ = v13
    _ = v14
    return ""
}

func main() {
    var v1 int32 = 1
    var v2 int = 128
    var v3 float64 = (((1.5 / 1e10) + 0.0) + (float64((3.14159 - 3.14159)) / 1e10))
    v4 := make(map[uint32]int)
    var v5 string = (("" + "a") + (("" + "\n\t") + "\n\t"))
    v6 := make(map[string]uint32)
    v15 := []int{(v2 & 128)}
    v16 := GenS1{F0: -32768}
    var v17 int64 = 0
    switch int(v17) {
    case 128:
        v17 += v17
    default:
        switch (^(v2)) {
        case 127:
            fmt.Println(v1, v1, v3)
        case 4294967295:
            for fi18 := 0; fi18 < 15; fi18++ {
                fmt.Println(v2, fi18, v3)
                var v19 uint = 625
                _ = fi18
                _ = v19
            }
        case 663:
            var v20 uint16 = 0
            _ = v20
        }
    }
    var v21 int64 = -862
    if false {
        for fi22 := 0; fi22 < 4; fi22++ {
            for fi23 := 0; fi23 < 10; fi23++ {
                var v24 int64 = int64(fi22)
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
        var v25 int32 = (v1 & v1)
        v1 = v25
        _ = v25
    } else {
        v2 -= v2
        var v26 string = v5
        var v27 string = v26
        _ = v26
        _ = v27
    }
    if true {
        var v28 int16 = 32767
        _ = v28
    }
    switch "" {
    case "hello":
        if len(v15) > 0 {
            v15[0] = (v2 >> ((0) & 63))
        }
    case "a":
        if false {
            v29 := GenS1{F0: -280}
            v30 := GenS1{F0: 0}
            var v31 int16 = -32768
            _ = v29
            _ = v30
            _ = v31
        } else {
            v32 := GenS1{F0: 1}
            _ = v32
        }
    case "foo bar":
        var v33 uint64 = 427
        _ = v33
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        var v35 int32 = v1
        if (("hello" + ("" + (("hello" + (v5 + v5)) + "hello"))) == v5) {
            for fi36 := 0; fi36 < 20; fi36++ {
                fmt.Println(v21)
                var v37 uint32 = 65535
                var v38 bool = true
                _ = fi36
                _ = v37
                _ = v38
            }
            delete(v6, v5)
            switch 518 {
            case 32768:
                v3 += (((v3 * v3) * 3.14159) + ((v3 + v3) + v3))
            }
        } else {
            var v39 uint = 65535
            if (true || false) {
                if true {
                    for fi40 := 0; fi40 < 16; fi40++ {
                        v5 = v5
                        _ = fi40
                    }
                    v39 = 255
                }
                var v41 int16 = -17
                if len(v15) > 0 {
                    v15[0] = fi34
                }
                _ = v41
            }
            _ = v39
        }
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
