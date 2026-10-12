package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint64, p1 uint16) int32 {
    switch 2147483647 {
    case -358:
        var v7 uint64 = 18446744073709551615
        _ = v7
    }
    var v8 uint8 = 1
    p0 -= 18446744073709551615
    p0 = (p0 | p0)
    _ = p0
    _ = p1
    _ = v8
    return 995
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    v4 := []int{127, -2147483648}
    v5 := []string{"", "a", "\n\t"}
    var v6 float64 = (-2.25 - -2.25)
    for fi9 := 0; fi9 < 13; fi9++ {
        if true {
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := GenS1{F0: 0, F1: 65535, F2: 0}
                v12 := GenS1{F0: 990, F1: 247, F2: 1}
                var v13 uint = 2
                v14 := GenS2{F0: 0}
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
                _ = v14
            }
            v15 := GenS2{F0: -450}
            _ = v15
        } else {
            v2[770] = ("a" + ((("\n\t" + ("hello" + "hello")) + "\n\t") + "hello"))
            v1 += 235
        }
        v6 = v6
        _ = fi9
    }
    var v16 int32 = 42
    switch -554 {
    case -957:
        v6 += ((1.5 * -2.25) + v6)
    case 2147483647:
        var v17 float32 = ((3.14159 * float32(v6)) * (3.14159 + -2.25))
        _ = v17
    case -962:
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 float64 = v6
            _ = fi18
            _ = v19
        }
    }
    fmt.Println(v6, v1)
    switch "hello" {
    case "":
        var v20 float32 = (0.1 / 1.0)
        _ = v20
    case "\n\t":
        v21 := GenS1{F0: 2, F1: 0, F2: 4294967295}
        _ = v21
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        v23 := GenS2{F0: -327}
        if (true || (42 >= 65535)) {
            if (false || false) {
                v3[fi22] = 4294967295
            } else {
                var v24 int = (-(fi22))
                _ = v24
            }
            for fi25 := 0; fi25 < 6; fi25++ {
                v26 := v3[fi25]
                _ = fi25
                _ = v26
            }
            switch "" {
            case "foo bar":
                v27 := GenS1{F0: 255, F1: 2, F2: 2}
                _ = v27
            }
        }
        var v28 float64 = v6
        _ = fi22
        _ = v23
        _ = v28
    }
    fmt.Println(v1, v1)
    for fi29 := 0; fi29 < 5; fi29++ {
        var v30 uint = 2
        v31 := GenS2{F0: 221}
        _ = fi29
        _ = v30
        _ = v31
    }
    fmt.Println(v6, v16)
    if true {
        for fi32 := 0; fi32 < 19; fi32++ {
            switch (fi32 - fi32) {
            case 1:
                v33 := GenS3{F0: -32768, F1: 2147483647}
                _ = v33
            }
            _ = fi32
        }
        v16 += v16
        var v34 uint32 = 264
        _ = v34
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
