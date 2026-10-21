package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 string) int16 {
    for fi5 := 0; fi5 < 12; fi5++ {
        p0 = ((p0 + p0) + "foo bar")
        var v6 uint64 = 18446744073709551615
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 42}
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            p0 = "\n\t"
            _ = fi9
        }
        var v10 float64 = (1e10 / 1e10)
        v10 += ((v10 - float64(0.1)) / (v10 / (float64(1.5) - v10)))
        _ = fi8
        _ = v10
    }
    var v11 uint8 = 0
    _ = p0
    _ = v11
    _ = v7
    return 32767
}

func main() {
    v1 := []int{-870, 255}
    v2 := make(map[string]string)
    v3 := make(map[string]bool)
    var v4 int8 = 115
    var v12 float32 = (float32(1.5) * (0.0 * (((-2.25 + -2.25) / float32((1.5 * -2.25))) + 1.5)))
    switch (("" + "") + "") {
    case "hello":
        v4 = (v4 + 100)
    }
    v12 = v12
    v12 -= 0.0
    switch 9223372036854775807 {
    case 0:
        v12 = 3.14159
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        switch "foo bar" {
        case "foo bar":
            var v14 uint8 = 255
            _ = v14
        case "hello":
            var v15 int8 = v4
            _ = v15
        }
        var v16 float64 = 0.0
        _ = fi13
        _ = v16
    }
    var v17 string = "hello"
    for fi18 := 0; fi18 < 14; fi18++ {
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 uint64 = 0
            _ = fi19
            _ = v20
        }
        switch (-(fi18)) {
        case -2147483648:
            v4 += v4
        }
        if false {
            var v21 uint = 4294967295
            var v22 int32 = 42
            fmt.Println(v17)
            _ = v21
            _ = v22
        } else {
            if true {
                if false {
                    v23 := v2[(v17 + v17)]
                    _ = v23
                } else {
                    if true {
                        v24 := GenS1{F0: -360}
                        var v25 int8 = (v4 >> ((65535) & 63))
                        _ = v24
                        _ = v25
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v17 = v17
                    }
                    var v26 float64 = 0.1
                    v17 = v17
                    _ = v26
                }
                var v27 string = v17
                _ = v27
            }
            if false {
                var v28 int32 = 2147483647
                _ = v28
            } else {
                v29 := v2["foo bar"]
                v29 = ((v29 + v29) + v29)
                _ = v29
            }
            v4 = (v4 - v4)
        }
        _ = fi18
    }
    v12 = v12
    if false {
        v17 = ("foo bar" + "")
    }
    switch (v17 + "") {
    case "":
        if len(v1) > 0 {
            v1[0] = -2147483648
        }
    }
    if true {
        v4 = (v4 ^ -128)
        var v30 uint32 = 986
        _ = v30
    } else {
        if true {
            for fi31 := 0; fi31 < 17; fi31++ {
                var v32 string = "hello"
                fmt.Println(v4, v17)
                v4 -= v4
                _ = fi31
                _ = v32
            }
            delete(v3, "hello")
            var v33 int32 = -1
            _ = v33
        }
    }
    delete(v3, v17)
    var v34 float32 = v12
    v17 = "\n\t"
    for fi35 := 0; fi35 < 11; fi35++ {
        v34 += v12
        if false {
            for fi36 := 0; fi36 < 7; fi36++ {
                v12 = ((v12 * (v12 * -2.25)) - v34)
                _ = fi36
            }
        } else {
            fmt.Println(v34, fi35)
        }
        _ = fi35
    }
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
