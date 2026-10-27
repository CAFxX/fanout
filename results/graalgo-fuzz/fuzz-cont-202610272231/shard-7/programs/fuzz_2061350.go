package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint8; F2 int }
type GenS3 struct { F0 bool; F1 int8; F2 bool }
type GenS4 struct { F0 uint8; F1 int8; F2 int64 }

func main() {
    v1 := []string{("" + ("" + "hello"))}
    var v2 float64 = (((0.0 - -2.25) + 1.5) * 0.0)
    v3 := make(map[string]bool)
    for fi4 := 0; fi4 < 12; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            v2 = v2
            v1 = append(v1, ("foo bar" + "hello"))
            var v6 float64 = float64(-2.25)
            _ = fi5
            _ = v6
        }
        if (!((-32768 > 1) || false)) {
            var v7 uint8 = 1
            if (1 < -49) {
                _ = len(v3)
                _ = len(v3)
            } else {
                v2 += (v2 - v2)
                v7 += v7
            }
            _ = v7
        } else {
            if (32767 > 617) {
                v2 = (float64(float32(v2)) / (v2 * (((0.0 / v2) * v2) / v2)))
            } else {
                var v8 string = "foo bar"
                v9 := GenS1{F0: 1}
                _ = v8
                _ = v9
            }
            for fi10 := 0; fi10 < 18; fi10++ {
                v11 := GenS2{F0: 65, F1: 128, F2: fi10}
                v2 = float64(3.14159)
                v3[("foo bar" + "hello")] = ((-(fi4)) == fi4)
                _ = fi10
                _ = v11
            }
            v2 -= (1.5 * v2)
        }
        _ = fi4
    }
    v2 -= v2
    v2 = (v2 * float64(0.1))
    if (0 > 1) {
        delete(v3, ("hello" + "a"))
    } else {
        v12 := v3[("hello" + "hello")]
        _ = v12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 int16 = 119
        fmt.Println(fi13)
        switch -39 {
        case -326:
            var v15 uint8 = 0
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    var v16 uint32 = 42
    fmt.Println(v16)
    var v17 uint32 = v16
    if false {
        fmt.Println(v16, v2)
        if ((v2 / float64((1e10 - 3.14159))) <= (1.5 * v2)) {
            var v18 uint8 = 187
            if (true || true) {
                if true {
                    fmt.Println(v18, v2)
                }
            }
            if true {
                v17 -= v17
            }
            _ = v18
        }
        var v19 int16 = 0
        _ = v19
    } else {
        _ = len(v1)
        v20 := GenS3{F0: (1 == -55), F1: 0, F2: true}
        if (-1 > 32767) {
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            fmt.Println(v16)
            v17 = v17
        } else {
            for fi21 := 0; fi21 < 1; fi21++ {
                delete(v3, "hello")
                _ = fi21
            }
            var v22 uint64 = 693
            _ = v22
        }
        _ = v20
    }
    if true {
        if len(v1) > 0 {
            v1[0] = ""
        }
        for fi23 := 0; fi23 < 18; fi23++ {
            for fi24 := 0; fi24 < 4; fi24++ {
                var v25 uint32 = v17
                v2 = 3.14159
                v16 -= (v17 >> ((1) & 63))
                var v26 uint8 = 123
                _ = fi24
                _ = v25
                _ = v26
            }
            switch ((("hello" + "hello") + "") + (("foo bar" + "") + "hello")) {
            case "hello":
                v27 := GenS4{F0: 1, F1: -1, F2: -567}
                _ = v27
            case "":
                var v28 string = "a"
                _ = v28
            }
            fmt.Println(fi23)
            _ = fi23
        }
    } else {
        var v29 int16 = 42
        _ = v29
    }
    v16 = (^(v17))
    v3["\n\t"] = (255 >= 32767)
    if false {
        if len(v1) > 0 {
            v1[0] = ("a" + ("hello" + "\n\t"))
        }
    } else {
        if (v2 < 1.5) {
            var v30 uint8 = 2
            _ = v30
        } else {
            v31 := GenS2{F0: -1, F1: 128, F2: -238}
            _ = v31
        }
    }
    fmt.Println(v16, v17, v2)
    for fi32 := 0; fi32 < 8; fi32++ {
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        v1 = append(v1, ("foo bar" + "hello"))
        var v33 string = "\n\t"
        _ = fi32
        _ = v33
    }
    v34 := GenS1{F0: 0}
    var v35 float32 = 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v35
}
