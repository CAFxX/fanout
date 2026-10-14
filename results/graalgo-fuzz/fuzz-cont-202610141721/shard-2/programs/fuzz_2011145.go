package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 string) uint32 {
    if (0 != 595) {
        var v6 uint64 = 18446744073709551615
        _ = v6
    } else {
        p0 = "hello"
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 uint16 = 255
            _ = fi7
            _ = v8
        }
        p0 = p0
    }
    if true {
        v9 := []int{32767, 127, 127, 1}
        var v10 string = ""
        _ = v10
        _ = v9
    }
    if (true || (0 <= 42)) {
        var v11 float32 = 3.14159
        _ = v11
    } else {
        p0 = p0
        var v12 uint16 = 32767
        _ = v12
    }
    if ((true && (32767 == 4294967295)) && true) {
        v13 := []int{-319}
        for fi14 := 0; fi14 < 7; fi14++ {
            p0 = "\n\t"
            _ = fi14
        }
        fmt.Println(p0)
        _ = v13
    } else {
        switch ((p0 + "foo bar") + p0) {
        case "":
            fmt.Println(p0)
        }
    }
    fmt.Println(p0)
    _ = p0
    return 150
}
func genFn1(p0 bool, p1 int64, p2 int8) uint16 {
    p1 += (^(p1))
    if ((!true) || false) {
        for fi15 := 0; fi15 < 11; fi15++ {
            p1 = p1
            _ = fi15
        }
    }
    var v16 uint64 = 299
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 0
}

func main() {
    v1 := []uint32{915}
    v2 := []uint32{2147483647}
    v3 := make(map[uint32]string)
    var v4 int64 = 1
    var v5 int = 127
    var v17 int64 = v4
    v5 -= 1
    if true {
        v4 = v4
        v18 := GenS1{F0: 42}
        var v19 bool = false
        _ = v18
        _ = v19
    } else {
        v4 -= v4
        if true {
            var v20 int32 = 1
            switch ("\n\t" + "hello") {
            case "hello":
                v5 -= (v5 + v5)
            case "\n\t":
                v1 = append(v1, 4294967295)
            case "foo bar":
                if false {
                    v21 := GenS1{F0: 0}
                    v3[42] = "foo bar"
                    v22 := GenS1{F0: 2}
                    _ = v21
                    _ = v22
                }
            }
            var v23 int8 = int8(v17)
            _ = v20
            _ = v23
        } else {
            v5 += 9223372036854775807
            for fi24 := 0; fi24 < 18; fi24++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi24
            }
            v25 := GenS1{F0: 0}
            _ = v25
        }
        var v26 int8 = 127
        _ = v26
    }
    var v27 uint = 817
    if (255 <= 1) {
        for fi28 := 0; fi28 < 3; fi28++ {
            v5 += fi28
            _ = fi28
        }
        if (!false) {
            v27 = (v27 * v27)
        } else {
            fmt.Println(v27, v4)
            var v29 float32 = 0.0
            _ = v29
        }
    } else {
        var v30 bool = false
        fmt.Println(v27, v17)
        v4 = (v17 << (((v27 >> ((v27) & 63))) & 63))
        _ = v30
    }
    v31 := GenS1{F0: 2}
    v5 = 32767
    _ = len(v1)
    v2 = append(v2, 199)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
