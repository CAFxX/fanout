package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    var v4 int8 = 42
    if false {
        for fi5 := 0; fi5 < 6; fi5++ {
            var v6 float64 = float64(3.14159)
            var v7 uint8 = 128
            _ = fi5
            _ = v6
            _ = v7
        }
    } else {
        v1 -= 18446744073709551615
        v4 -= (v4 % ((v4) | 1))
        var v8 float32 = -2.25
        _ = v8
    }
    var v9 int16 = 32767
    for fi10 := 0; fi10 < 2; fi10++ {
        if false {
            var v11 int8 = v4
            for fi12 := 0; fi12 < 13; fi12++ {
                v4 += v11
                _ = fi12
            }
            for fi13 := 0; fi13 < 3; fi13++ {
                v14 := GenS1{F0: 125}
                _ = fi13
                _ = v14
            }
            _ = v11
        }
        _ = fi10
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        switch fi15 {
        case 4294967295:
            v1 = v1
        case -9223372036854775808:
            v4 -= v4
        case -2147483648:
            delete(v2, (("foo bar" + "") + ""))
        }
        _ = fi15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        v4 += 0
        var v17 uint16 = 32767
        v9 = (v9 - -1)
        _ = fi16
        _ = v17
    }
    v18 := []int{}
    if (!true) {
        v9 -= v9
        v19 := GenS1{F0: 127}
        _ = v19
    } else {
        for fi20 := 0; fi20 < 2; fi20++ {
            v4 += (-(v4))
            if len(v18) > 0 {
                v18[0] = -2147483648
            }
            fmt.Println(v4, v4, v4)
            _ = fi20
        }
        v21 := GenS1{F0: 42}
        _ = v21
    }
    v2[""] = "hello"
    v22 := GenS1{F0: 128}
    for fi23 := 0; fi23 < 12; fi23++ {
        v2["foo bar"] = "hello"
        if false {
            var v24 string = "foo bar"
            v25 := GenS1{F0: 2}
            _ = v24
            _ = v25
        } else {
            v1 = 255
        }
        _ = fi23
    }
    delete(v3, 2147483647)
    v1 = 18446744073709551615
    switch ("\n\t" + "") {
    case "":
        delete(v2, "")
    }
    v26 := GenS1{F0: 1}
    v4 += int8(v1)
    fmt.Println(v4, v9)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v9
}
