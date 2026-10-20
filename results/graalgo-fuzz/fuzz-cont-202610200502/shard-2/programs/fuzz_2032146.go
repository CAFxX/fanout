package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 0
    var v3 int64 = 1
    var v4 int8 = -128
    for fi5 := 0; fi5 < 14; fi5++ {
        delete(v1, 401)
        var v6 uint = 103
        if (0 > v3) {
            if (("hello" + "\n\t") != ("a" + "hello")) {
                v2 = v2
            }
            switch "hello" {
            case "foo bar":
                v3 -= v3
            }
        }
        _ = fi5
        _ = v6
    }
    switch (("\n\t" + "a") + "\n\t") {
    case "foo bar":
        if (174 != 698) {
            v7 := []int{}
            _ = v7
        } else {
            var v8 int8 = -128
            switch 127 {
            case -2147483648:
                for fi9 := 0; fi9 < 9; fi9++ {
                    fmt.Println(v4, v3, fi9)
                    _ = fi9
                }
            default:
                fmt.Println(v3)
            }
            _ = v8
        }
    }
    var v10 int64 = v3
    switch (("hello" + "hello") + "hello") {
    case "a":
        v11 := GenS1{F0: (-(v2))}
        _ = v11
    }
    v12 := GenS1{F0: (v2 << ((113) & 63))}
    v13 := GenS1{F0: 2147483647}
    var v14 float32 = 1.5
    for fi15 := 0; fi15 < 15; fi15++ {
        for fi16 := 0; fi16 < 5; fi16++ {
            v17 := v1[4294967295]
            v18 := []int{fi15}
            v1[284] = v17
            _ = fi16
            _ = v17
            _ = v18
        }
        v19 := []int{fi15, (fi15 % (((fi15 ^ -2147483648)) | 1))}
        var v20 int64 = (v10 >> ((77) & 63))
        _ = fi15
        _ = v19
        _ = v20
    }
    switch "" {
    case "hello":
        switch "a" {
        case "a":
            v21 := v1[2147483648]
            _ = v21
        default:
            var v22 uint = 968
            _ = v22
        }
    case "\n\t":
        v23 := v1[709]
        _ = v23
    default:
        v2 = (v2 % ((v2) | 1))
    }
    if false {
        v24 := GenS1{F0: v2}
        v25 := GenS1{F0: (v2 * v2)}
        var v26 uint32 = 967
        _ = v24
        _ = v25
        _ = v26
    } else {
        for fi27 := 0; fi27 < 6; fi27++ {
            switch "a" {
            case "\n\t":
                v28 := GenS1{F0: v2}
                _ = v28
            }
            for fi29 := 0; fi29 < 16; fi29++ {
                var v30 float64 = ((1.5 / 0.1) / 0.1)
                var v31 float64 = v30
                _ = fi29
                _ = v30
                _ = v31
            }
            v10 = v10
            _ = fi27
        }
        var v32 uint8 = 42
        v33 := []int{0, -781, -965}
        _ = v32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
