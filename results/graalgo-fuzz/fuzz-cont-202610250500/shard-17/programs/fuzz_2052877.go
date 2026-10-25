package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
func genFn0() bool {
    v5 := GenS1{F0: 450, F1: 1}
    switch "a" {
    case "a":
        v6 := []int{9223372036854775807, 32768, 4294967295}
        _ = v6
    }
    v7 := GenS1{F0: 0, F1: 42}
    if false {
        v8 := GenS1{F0: 1, F1: 42}
        var v9 float64 = -2.25
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 int16 = -32768
        v11 += v11
        _ = fi10
        _ = v11
    }
    _ = v5
    _ = v7
    return false
}
func genFn1(p0 int8) uint64 {
    p0 = p0
    var v12 string = "a"
    var v13 int = -1
    _ = p0
    _ = v12
    _ = v13
    return 393
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 65535
    var v3 string = "hello"
    var v4 int16 = 1
    if (!(false || (220 <= 127))) {
        v3 = v3
        if false {
            if (!(!true)) {
                switch 2147483647 {
                case 42:
                    if true {
                        v14 := v1[2]
                        _ = v14
                    }
                }
                v3 = ("a" + ((((((v3 + "hello") + "hello") + "\n\t") + v3) + v3) + v3))
                var v15 bool = true
                _ = v15
            } else {
                var v16 int8 = 127
                var v17 uint16 = 100
                _ = v16
                _ = v17
            }
        }
        v4 = v4
    }
    var v18 uint64 = 0
    var v19 bool = (-2.25 >= -2.25)
    var v20 int64 = 9223372036854775807
    switch -994 {
    case 128:
        v21 := GenS1{F0: -1, F1: genFn1(95)}
        _ = v21
    case -13:
        v20 = v20
    }
    switch v3 {
    case "":
        for fi22 := 0; fi22 < 1; fi22++ {
            switch fi22 {
            case 32767:
                v3 = v3
            case 646:
                _ = len(v1)
            }
            _ = fi22
        }
    default:
        _ = len(v1)
    }
    v2 = (^(v2))
    var v23 int32 = -1
    var v24 int = 42
    if (((v2 >> ((v2) & 63)) >= v2) || ((3.14159 / 1.0) >= 3.14159)) {
        v25 := GenS2{F0: v24}
        var v26 int8 = 49
        _ = v25
        _ = v26
    } else {
        v4 = -1
        for fi27 := 0; fi27 < 15; fi27++ {
            v24 = v24
            _ = fi27
        }
    }
    var v28 int = -2147483648
    v18 = genFn1(100)
    if true {
        var v29 int64 = v20
        for fi30 := 0; fi30 < 11; fi30++ {
            var v31 float64 = float64((0.0 + -2.25))
            v29 = v20
            _ = fi30
            _ = v31
        }
        var v32 float64 = -2.25
        _ = v29
        _ = v32
    } else {
        var v33 float32 = -2.25
        v2 -= v2
        if true {
            v33 = v33
        }
        _ = v33
    }
    v34 := v1[2147483647]
    for fi35 := 0; fi35 < 3; fi35++ {
        for fi36 := 0; fi36 < 10; fi36++ {
            v1[42] = (v19 || false)
            v37 := v1[1]
            var v38 float64 = ((0.1 / 1e10) / ((-2.25 / (0.1 * float64(1.5))) * float64((1.5 / -2.25))))
            if v34 {
                v34 = (0.1 == 3.14159)
                v34 = ((42 <= 255) && (!(2147483648 > 697)))
                fmt.Println(v37, v20)
                fmt.Println(v34)
                v39 := []int{fi36, -1, v28, -1}
                _ = v39
            }
            _ = fi36
            _ = v37
            _ = v38
        }
        v19 = v19
        for fi40 := 0; fi40 < 17; fi40++ {
            v3 = "\n\t"
            _ = fi40
        }
        _ = fi35
    }
    switch (-(v28)) {
    case 127:
        for fi41 := 0; fi41 < 1; fi41++ {
            v34 = true
            if (117 != 127) {
                v4 += v4
                var v42 uint = (v2 / ((v2) | 1))
                _ = v42
            }
            _ = fi41
        }
    }
    v43 := GenS3{F0: v20}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v43
}
